# CDN Ranges

Tool to download a list of IPv4 and IPv6 ranges used by CDNs. This helps to avoid performing unnecessary port scans when 
doing bug bounties.

This uses publicly available lists of IP ranges, provided by most providers, and [IP Guide](https://ip.guide) to query IP ranges for ASNs.

## CDN Providers

| Provider   | ASN or Public List                                                                                                                                                                                                                                                                                                                                                                          |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Akamai     | AS12222, AS16625, AS16702, AS17204, AS18680, AS18717, AS20189, AS20940, AS21342, AS21357, AS21399,  AS22207,  AS22452,  AS23454,  AS23455,  AS23903,  AS24319,  AS26008,  AS30675,  AS31107,  AS31108,  AS31109,  AS31110,  AS31377,  AS33047,  AS33905,  AS34164,  AS34850,  AS35204,  AS35993,  AS35994,  AS36183,  AS39836,  AS43639,  AS55409,  AS55770,  AS63949,  AS133103,  AS393560 |
| ArvanCloud | https://www.arvancloud.ir/en/ips.txt                                                                                                                                                                                                                                                                                                                                                        |
| Bunny      | AS200325                                                                                                                                                                                                                                                                                                                                                                                    |
| CacheFly   | https://cachefly.cachefly.net/ips/rproxy.txt                                                                                                                                                                                                                                                                                                                                                |
| CDN77      | https://prefixlists.tools.cdn77.com/public_lmax_prefixes.json                                                                                                                                                                                                                                                                                                                               |
| CDNetworks | AS36408                                                                                                                                                                                                                                                                                                                                                                                     |
| Cloudflare | https://www.cloudflare.com/ips-v4 https://www.cloudflare.com/ips-v6                                                                                                                                                                                                                                                                                                                         |
| CloudFront | https://ip-ranges.amazonaws.com/ip-ranges.json                                                                                                                                                                                                                                                                                                                                              |
| DDoS-Guard | AS57724                                                                                                                                                                                                                                                                                                                                                                                     |
| Edgecast   | AS15133                                                                                                                                                                                                                                                                                                                                                                                     |
| EdgeNext   | AS139057, AS149981                                                                                                                                                                                                                                                                                                                                                                          |
| Edgio      | AS60261                                                                                                                                                                                                                                                                                                                                                                                     |
| F5         | https://docs.cloud.f5.com/docs-v2/platform/reference/network-cloud-ref                                                                                                                                                                                                                                                                                                                      |
| Fastly     | https://api.fastly.com/public-ip-list                                                                                                                                                                                                                                                                                                                                                       |
| Gcore      | https://api.gcore.com/cdn/public-ip-list                                                                                                                                                                                                                                                                                                                                                    |
| Imperva    | https://my.imperva.com/api/integration/v1/ips                                                                                                                                                                                                                                                                                                                                               |
| Leaseweb   | https://networksdb.io/ip-addresses-of/leaseweb-cdn-bv                                                                                                                                                                                                                                                                                                                                       |
| Limelight  | AS22822                                                                                                                                                                                                                                                                                                                                                                                     |
| Medianova  | https://cloud.medianova.com/api/v1/ip/blocks-list                                                                                                                                                                                                                                                                                                                                           |
| Qrator     | AS200449                                                                                                                                                                                                                                                                                                                                                                                    |
| StackPath  | AS12989                                                                                                                                                                                                                                                                                                                                                                                     |
| StormWall  | AS59796                                                                                                                                                                                                                                                                                                                                                                                     |
| Sucuri     | AS30148                                                                                                                                                                                                                                                                                                                                                                                     |
| X4B        | AS136165                                                                                                                                                                                                                                                                                                                                                                                    |

If a provider is missing, please open an issue with a link to their IP ranges or ASN.

## Installation

Go install

```bash
go install -v github.com/taythebot/cdn-ranges/cmd/cdn-ranges@latest
```

Manual Build

```bash
git clone https://github.com/taythebot/cdn-ranges
go build -o cdn-ranges cmd/cdn-ranges.main.go 
```


## Usage

Download IPv4 and IPv6 ranges for all providers

```bash
cdn-ranges -output ranges.txt
```

Download IPV4 ranges only

```bash
cdn-ranges -ipv4 -output ranges.txt
```

Download IPV6 ranges only

```bash
cdn-ranges -ipv6 -output ranges.txt
```

Download for a specific provider (lowercase)

```bash
cdn-ranges -provider cloudflare
```

Dump in json format

```bash
cdn-ranges -format json -output ranges.json
```

Dump in csv format (provider,type,range)

```bash
cdn-ranges -format csv -output ranges.csv
```

## Output Formats

* txt (default)
* json
* csv
