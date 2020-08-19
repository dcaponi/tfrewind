# Terraform Rewind

#### Back-parse your tfstate into a useable main.tf plan

## What Is It?

`tfrewind` is a tool that can back-convert your tfstate file into a useable main.tf.

This is especially useful if you are running `terraform import` and want to convert the resulting `.tfstate` file into something immediately useable

## Getting Started

**From Source - requires go**

1. clone this repository
2. from this repository run `go install .`
3. from the folder where a `.tfstate` file lives, call `tfrewind`

**From Binary**

1. grab the binary for your OS from this release and add it to your `$gopath/bin` directory or wherever your PATH looks for binaries

e.g. on mac its usually `/usr/bin/` or `/usr/local/bin` as long as that folder is on your $PATH

## Example
An example `terraform.tfstate` file has been given to demonstrate what will happen. If you installed the app correctly, running `tfrewind` from this directory with the provided example, you should get an output `main.tf` that looks like the provided `main_example.tf`.

If you were to `terraform import` your own resources prior to running this, you'd get a proper `main.tf`.

## Contribute

I started this thinking it'd be cool to add some functionality to my Terraform workflow and as such I haven't added everything under the sun, but if there's a resource you need imported, please add it.

1. You'll need to add a file to `/resources` with the convention `provider_resource.go` so for example `aws_iam_user.go` or `onelogin_user.go`
2. In that file implement the `HCLShape() interface{}` method. You'll return the address to a struct that will be used for marshalling/unmarshalling json from the `.tfstate` file. (See `aws_iam_user.go` for example)
3. Add the parseable resource you just set up to the `parsables-factory.go` file
```
case "<your_parsable>":
  imf.parsables[parsableType] = &<your_parsable>{}
```

## Supported Terraform Versions
Currently supports only `v0.12.xx`

## Planned Updates
Support Terraform `v0.13.xx`.
Automatic imports based on providers' `list resources` APIs 
