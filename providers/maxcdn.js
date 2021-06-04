const axios = require('axios')

module.exports = async () => {
  const { data } = await axios.get(
    'https://support.maxcdn.com/hc/en-us/article_attachments/360051920551/maxcdn_ips.txt'
  )

  return data.split('\n').slice(0, -1)
}
