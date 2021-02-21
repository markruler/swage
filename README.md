# Swage

> Command line tool to convert OpenAPI specification data (`json`, `yaml`) to Excel (`xlsx`) format

## Usage

```bash
swage gen <path>
```

### Example

```bash
swage gen examples/testdata/yaml/docker.v1.41.yaml
```

```bash
swage gen https://docs.docker.com/engine/api/v1.41.yaml
```

```bash
swage gen https://raw.githubusercontent.com/kubernetes/kubernetes/master/api/openapi-spec/swagger.json
```

## Background

> OAS
>
> - [Swagger - SmartBear Software](https://swagger.io/docs/specification/about)
> - [OpenAPI - The Linux Foundation](https://www.openapis.org/about)
> - OpenAPI Specification (formerly Swagger Specification) is an API description format for REST APIs. An OpenAPI file allows you to describe your entire API.

> [XLSX extension](https://docs.microsoft.com/en-us/openspecs/office_standards/ms-xlsx/)
>
> - The Excel (.xlsx) Extensions to the Office Open XML SpreadsheetML File Format specifies extensions
>   to the Office Open XML file formats described in [ISO/IEC29500-1:2016](https://www.iso.org/standard/71691.html).
>   The extensions are specified using conventions provided by the Office Open XML file formats
>   described in [ISO/IEC29500-3:2015](https://www.iso.org/standard/65533.html).

## Dependencies

- [spf13/cobra](https://github.com/spf13/cobra)
- [go-openapi/spec](https://github.com/go-openapi/spec)
- [go-openapi/loads](https://github.com/go-openapi/loads)
- [360EntSecGroup-Skylar/excelize](https://github.com/360EntSecGroup-Skylar/excelize)

## References

- [SmartBear Documentations](https://swagger.io/docs/specification)
- [OAI OpenAPI Specification](https://github.com/OAI/OpenAPI-Specification)
- [APIs.guru](https://apis.guru/browse-apis/)
- [APIs.guru - GitHub](https://github.com/APIs-guru/openapi-directory)

## Open API Specification (OAS) Revision History

- [OpenAPI Specification - Appendix A: Revision History](https://swagger.io/specification/#appendix-a-revision-history)
- by SmartBear

| Version   | Date       | Notes                                             |
| --------- | ---------- | ------------------------------------------------- |
| 3.0.3     | 2020-02-20 | Patch release of the OpenAPI Specification 3.0.3  |
| 3.0.2     | 2018-10-08 | Patch release of the OpenAPI Specification 3.0.2  |
| 3.0.1     | 2017-12-06 | Patch release of the OpenAPI Specification 3.0.1  |
| 3.0.0     | 2017-07-26 | Release of the OpenAPI Specification 3.0.0        |
| 3.0.0-rc2 | 2017-06-16 | rc2 of the 3.0 specification                      |
| 3.0.0-rc1 | 2017-04-27 | rc1 of the 3.0 specification                      |
| 3.0.0-rc0 | 2017-02-28 | Implementer's Draft of the 3.0 specification      |
| 2.0       | 2015-12-31 | Donation of Swagger 2.0 to the OpenAPI Initiative |
| 2.0       | 2014-09-08 | Release of Swagger 2.0                            |
| 1.2       | 2014-03-14 | Initial release of the formal document.           |
| 1.1       | 2012-08-22 | Release of Swagger 1.1                            |
| 1.0       | 2011-08-10 | First release of the Swagger Specification        |

![OAS Version](./oas-version.jpg)

_[A Guide to What’s New in OpenAPI 3.0](https://swagger.io/blog/news/whats-new-in-openapi-3-0/) - Ryan Pinkham_
