const axios = require('axios')

module.exports = async () => {
  const { data } = await axios.get(
    'https://ip-ranges.amazonaws.com/ip-ranges.json'
  )

  return data.prefixes.filter(({ service }) => service === 'CLOUDFRONT')
}
