# EInvoice

## OpenAPI 3 / Swagger

[OpenAPI 3 definition](docs/swagger.yml).

You can view it on this [website](https://generator.swagger.io/?url=https://raw.githubusercontent.com/filipsladek/einvoice/master/docs/swagger.yml).

## Try it out

[https://web-app.dev.filipsladek.com](https://web-app.dev.filipsladek.com)

user: E0000046137

password: PopradTa3@

## Docs

[Einvoice - API - discussion](https://docs.google.com/document/d/1gjDwwu6qCBvdB63i9mKdPWPGL-Y76UJ2LWby6kaNKoA/edit?usp=sharing)

[eFakturacia - pohlad z trhu](https://docs.google.com/document/d/1RdCJ-erer9rOD41Tnc9JLMCZw2fXWkUnO-E5DM4BFu0/edit?usp=sharing)

## Discussion

[platforma.slovensko.digital](https://platforma.slovensko.digital/t/red-flags-informacny-system-elektronickej-fakturacie-is-efa/5640/83?u=filip_sladek)

## Development

* Ensure `postgres` and `redis-server` services are running.
* Initialize DB (TODO: replace by migrations)

    `./db.sh --drop --create --setup`

* Export proper env variables for every service in `${server-name}/.env`.
You can initialize it by `cp .env-template .env` and update for your use.

* Run services you need:

```shell script
./dev-scripts/start_service.sh ${service-name}
```

* Finally run [web-app](einvoice-web-app/README.md)

## Deployment

[Ansible](ansible/README.md)

## XML

### UBL2.1

[Maindoc](http://docs.oasis-open.org/ubl/os-UBL-2.1/xsd)

From maindoc you need only [this](http://docs.oasis-open.org/ubl/os-UBL-2.1/xsd/maindoc/UBL-Invoice-2.1.xsd)
part.

### D16B

Subset for [CrossIndustryInvoice](https://www.unece.org/fileadmin/DAM/cefact/xml_schemas/D16B_SCRDM__Subset__CII.zip)
