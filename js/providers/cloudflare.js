const axios = require('axios')

module.exports = async () => {
  const { data } = await axios.get('https://www.cloudflare.com/ips-v4')

  return data.split('\n').slice(0, -1)
}
