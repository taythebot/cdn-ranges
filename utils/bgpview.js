const axios = require('axios')

module.exports = async (asn) => {
  const { data } = await axios.get(`https://api.bgpview.io/asn/${asn}/prefixes`)

  return data.data.ipv4_prefixes.map(({ prefix }) => prefix)
}
