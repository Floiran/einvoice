# EInvoice

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

From maindoc you need only [this](http://docs.oasis-open.org/ubl/os-UBL-2.1/xsd/maindoc/UBL-Invoice-2.1.xsd) part.

### D16B

Subset for [CrossIndustryInvoice](https://www.unece.org/fileadmin/DAM/cefact/xml_schemas/D16B_SCRDM__Subset__CII.zip)
