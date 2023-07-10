const axios = require('axios')

module.exports = async () => {
  const { data } = await axios.get(
    'https://ip-ranges.amazonaws.com/ip-ranges.json'
  )

  return data.prefixes.reduce((arr, { service, ip_prefix }) => {
    if (service === 'CLOUDFRONT') arr.push(ip_prefix)
    return arr
  }, [])
}
