const axios = require('axios')

module.exports = async () => {
  const { data } = await axios.get('https://api.fastly.com/public-ip-list')

  return data.addresses
}
