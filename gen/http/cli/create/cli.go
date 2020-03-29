// Code generated by goa v3.1.1, DO NOT EDIT.
//
// create HTTP client CLI support package
//
// Command:
// $ goa gen crud/design

package cli

import (
	createc "crud/gen/http/create/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `create (create|list|remove|update)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` create create --body '{
      "comments": [
         "Quasi ducimus.",
         "In ea iste.",
         "Provident atque nesciunt alias aspernatur."
      ],
      "id": 4071105877,
      "name": "or1"
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		createFlags = flag.NewFlagSet("create", flag.ContinueOnError)

		createCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		createCreateBodyFlag = createCreateFlags.String("body", "REQUIRED", "")

		createListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		createRemoveFlags  = flag.NewFlagSet("remove", flag.ExitOnError)
		createRemoveIDFlag = createRemoveFlags.String("id", "REQUIRED", "ID of blog to remove")

		createUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		createUpdateBodyFlag = createUpdateFlags.String("body", "REQUIRED", "")
		createUpdateIDFlag   = createUpdateFlags.String("id", "REQUIRED", "ID of blog to be updated")
	)
	createFlags.Usage = createUsage
	createCreateFlags.Usage = createCreateUsage
	createListFlags.Usage = createListUsage
	createRemoveFlags.Usage = createRemoveUsage
	createUpdateFlags.Usage = createUpdateUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "create":
			svcf = createFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "create":
			switch epn {
			case "create":
				epf = createCreateFlags

			case "list":
				epf = createListFlags

			case "remove":
				epf = createRemoveFlags

			case "update":
				epf = createUpdateFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "create":
			c := createc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = createc.BuildCreatePayload(*createCreateBodyFlag)
			case "list":
				endpoint = c.List()
				data = nil
			case "remove":
				endpoint = c.Remove()
				data, err = createc.BuildRemovePayload(*createRemoveIDFlag)
			case "update":
				endpoint = c.Update()
				data, err = createc.BuildUpdatePayload(*createUpdateBodyFlag, *createUpdateIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// createUsage displays the usage of the create command and its subcommands.
func createUsage() {
	fmt.Fprintf(os.Stderr, `The blog service gives blog details.
Usage:
    %s [globalflags] create COMMAND [flags]

COMMAND:
    create: Add new blog and return its ID.
    list: List all entries
    remove: Remove blog from storage
    update: Updating the existing blog

Additional help:
    %s create COMMAND --help
`, os.Args[0], os.Args[0])
}
func createCreateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] create create -body JSON

Add new blog and return its ID.
    -body JSON: 

Example:
    `+os.Args[0]+` create create --body '{
      "comments": [
         "Quasi ducimus.",
         "In ea iste.",
         "Provident atque nesciunt alias aspernatur."
      ],
      "id": 4071105877,
      "name": "or1"
   }'
`, os.Args[0])
}

func createListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] create list

List all entries

Example:
    `+os.Args[0]+` create list
`, os.Args[0])
}

func createRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] create remove -id UINT32

Remove blog from storage
    -id UINT32: ID of blog to remove

Example:
    `+os.Args[0]+` create remove --id 2634915235
`, os.Args[0])
}

func createUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] create update -body JSON -id UINT32

Updating the existing blog
    -body JSON: 
    -id UINT32: ID of blog to be updated

Example:
    `+os.Args[0]+` create update --body '{
      "comments": [
         "Quam sed nihil.",
         "Exercitationem eligendi.",
         "Laboriosam qui eaque eligendi facere pariatur.",
         "Facere qui minus doloremque beatae."
      ],
      "name": "Ea enim accusamus magni."
   }' --id 1699121745
`, os.Args[0])
}
