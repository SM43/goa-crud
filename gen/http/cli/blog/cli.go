// Code generated by goa v3.1.1, DO NOT EDIT.
//
// blog HTTP client CLI support package
//
// Command:
// $ goa gen github.com/sm43/goa-crud/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	blogc "github.com/sm43/goa-crud/gen/http/blog/client"
	oauthc "github.com/sm43/goa-crud/gen/http/oauth/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `oauth oauth
blog (create|list|show|remove|add)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` oauth oauth --body '{
      "token": "Aliquid provident."
   }'` + "\n" +
		os.Args[0] + ` blog create --body '{
      "blog": {
         "comments": [
            {
               "comment": "Eligendi quam eveniet non eaque omnis et.",
               "id": 5853931440448808029
            },
            {
               "comment": "Eligendi quam eveniet non eaque omnis et.",
               "id": 5853931440448808029
            }
         ],
         "name": "Soluta aut dolorum fuga rerum et et."
      }
   }' --auth "Ea ratione at et."` + "\n" +
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
		oauthFlags = flag.NewFlagSet("oauth", flag.ContinueOnError)

		oauthOauthFlags    = flag.NewFlagSet("oauth", flag.ExitOnError)
		oauthOauthBodyFlag = oauthOauthFlags.String("body", "REQUIRED", "")

		blogFlags = flag.NewFlagSet("blog", flag.ContinueOnError)

		blogCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		blogCreateBodyFlag = blogCreateFlags.String("body", "REQUIRED", "")
		blogCreateAuthFlag = blogCreateFlags.String("auth", "REQUIRED", "")

		blogListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		blogShowFlags  = flag.NewFlagSet("show", flag.ExitOnError)
		blogShowIDFlag = blogShowFlags.String("id", "REQUIRED", "ID of the blog to be fetched")

		blogRemoveFlags  = flag.NewFlagSet("remove", flag.ExitOnError)
		blogRemoveIDFlag = blogRemoveFlags.String("id", "REQUIRED", "ID of blog to remove")

		blogAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		blogAddBodyFlag = blogAddFlags.String("body", "REQUIRED", "")
		blogAddIDFlag   = blogAddFlags.String("id", "REQUIRED", "Id of the blog")
	)
	oauthFlags.Usage = oauthUsage
	oauthOauthFlags.Usage = oauthOauthUsage

	blogFlags.Usage = blogUsage
	blogCreateFlags.Usage = blogCreateUsage
	blogListFlags.Usage = blogListUsage
	blogShowFlags.Usage = blogShowUsage
	blogRemoveFlags.Usage = blogRemoveUsage
	blogAddFlags.Usage = blogAddUsage

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
		case "oauth":
			svcf = oauthFlags
		case "blog":
			svcf = blogFlags
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
		case "oauth":
			switch epn {
			case "oauth":
				epf = oauthOauthFlags

			}

		case "blog":
			switch epn {
			case "create":
				epf = blogCreateFlags

			case "list":
				epf = blogListFlags

			case "show":
				epf = blogShowFlags

			case "remove":
				epf = blogRemoveFlags

			case "add":
				epf = blogAddFlags

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
		case "oauth":
			c := oauthc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "oauth":
				endpoint = c.Oauth()
				data, err = oauthc.BuildOauthPayload(*oauthOauthBodyFlag)
			}
		case "blog":
			c := blogc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = blogc.BuildCreatePayload(*blogCreateBodyFlag, *blogCreateAuthFlag)
			case "list":
				endpoint = c.List()
				data = nil
			case "show":
				endpoint = c.Show()
				data, err = blogc.BuildShowPayload(*blogShowIDFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = blogc.BuildRemovePayload(*blogRemoveIDFlag)
			case "add":
				endpoint = c.Add()
				data, err = blogc.BuildAddPayload(*blogAddBodyFlag, *blogAddIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// oauthUsage displays the usage of the oauth command and its subcommands.
func oauthUsage() {
	fmt.Fprintf(os.Stderr, `The oauth service authorise user to access other APIs
Usage:
    %s [globalflags] oauth COMMAND [flags]

COMMAND:
    oauth: Github authentication to post a new blog

Additional help:
    %s oauth COMMAND --help
`, os.Args[0], os.Args[0])
}
func oauthOauthUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] oauth oauth -body JSON

Github authentication to post a new blog
    -body JSON: 

Example:
    `+os.Args[0]+` oauth oauth --body '{
      "token": "Aliquid provident."
   }'
`, os.Args[0])
}

// blogUsage displays the usage of the blog command and its subcommands.
func blogUsage() {
	fmt.Fprintf(os.Stderr, `The blog service gives blog details.
Usage:
    %s [globalflags] blog COMMAND [flags]

COMMAND:
    create: Add a new blog
    list: List all the blogs
    show: Show blog based on the id given
    remove: Delete a blog
    add: Add a new comment for a blog

Additional help:
    %s blog COMMAND --help
`, os.Args[0], os.Args[0])
}
func blogCreateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] blog create -body JSON -auth STRING

Add a new blog
    -body JSON: 
    -auth STRING: 

Example:
    `+os.Args[0]+` blog create --body '{
      "blog": {
         "comments": [
            {
               "comment": "Eligendi quam eveniet non eaque omnis et.",
               "id": 5853931440448808029
            },
            {
               "comment": "Eligendi quam eveniet non eaque omnis et.",
               "id": 5853931440448808029
            }
         ],
         "name": "Soluta aut dolorum fuga rerum et et."
      }
   }' --auth "Ea ratione at et."
`, os.Args[0])
}

func blogListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] blog list

List all the blogs

Example:
    `+os.Args[0]+` blog list
`, os.Args[0])
}

func blogShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] blog show -id UINT

Show blog based on the id given
    -id UINT: ID of the blog to be fetched

Example:
    `+os.Args[0]+` blog show --id 2757763773622452166
`, os.Args[0])
}

func blogRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] blog remove -id UINT

Delete a blog
    -id UINT: ID of blog to remove

Example:
    `+os.Args[0]+` blog remove --id 16522289313636908332
`, os.Args[0])
}

func blogAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] blog add -body JSON -id UINT

Add a new comment for a blog
    -body JSON: 
    -id UINT: Id of the blog

Example:
    `+os.Args[0]+` blog add --body '{
      "comments": {
         "comment": "Eligendi quam eveniet non eaque omnis et.",
         "id": 5853931440448808029
      }
   }' --id 2169912030633346718
`, os.Args[0])
}
