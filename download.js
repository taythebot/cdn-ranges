#!/usr/bin/env node

const { program } = require('commander')
const Listr = require('listr')
const { readdirSync, writeFileSync } = require('fs')

const formats = ['txt', 'csv', 'json']

const opts = program
  .description(
    'A tool to download IP ranges of CDN for bug bounties\nCreated by Tay (https://github.com/taythebot)'
  )
  .version('1.0.0')
  .requiredOption(
    '-f, --format <type>',
    'output format (txt, csv, json)',
    'txt'
  )
  .requiredOption('-o --output <file>', 'output file', 'ranges.txt')
  .option('-p --provider <names...>', 'download a specific provider')
  .parse()
  .opts()

;(async () => {
  if (!formats.includes(opts.format)) {
    return console.error(`Must be on of ${formats.join(', ')}`)
  }

  let providers = readdirSync(`${__dirname}/providers`).map(
    (files) => files.split('.js')[0]
  )

  if (opts.provider) {
    const invalidProvider = opts.provider.find(
      (provider) => !providers.includes(provider)
    )
    if (invalidProvider) {
      return console.error(
        `error: '${invalidProvider}' is not a valid provider`
      )
    }

    providers = Array.from(new Set(opts.provider))
  }

  const results = []
  const tasks = new Listr(
    providers.map((provider) => ({
      title: provider,
      task: async () => {
        const ranges = await require(`${__dirname}/providers/${provider}.js`)()
        ranges.forEach((range) => results.push({ provider, range }))
      },
    }))
  )

  console.log(`Fetching IP ranges from ${providers.length} providers:`)
  await tasks.run()

  let data
  switch (opts.format) {
    case 'csv':
      data = results
        .map(({ provider, range }) => `${provider},${range}`)
        .join('\n')
      break
    case 'json':
      data = JSON.stringify(results, null, 2)
      break
    case 'txt':
      data = results.map(({ range }) => range).join('\n')
      break
  }

  writeFileSync(opts.output, data, 'utf-8')
  console.log(`Successfully saved ${results.length} ranges in ${opts.output}`)
})()
