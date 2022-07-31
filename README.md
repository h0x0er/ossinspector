# ossinspector
A policy based open source package inspector 


## Build

```bash

  git clone https://github.com/h0x0er/ossinspector.git
  cd ossinspector
  cd main
  go build inspector.go -o ossinspector
  ./ossinspector
 
```

## Usage
  
  * Grab latest build from [here](https://github.com/h0x0er/ossinspector/releases)
  * Extract the file.
  * Copy ossinspector to `/usr/loca/sbin`

  1. Repo Mode

  * Create a `policy.yml` file
  * Run the `ossinspector`
  ```bash 
    ossinspector -p <location_to_policy_yml> -r owner/repo -rt score
  ```
  2. Project Mode
  * Create `policy.yml` file in your project folder
  * Run inspector with below command in project mode
  ```bash 
  ossinspector -t TYPE
  ```
  > TYPE: [node, python, go...]
  > Currently only node is supported
  


## policy.yml

```yml
# sample policy.yml
# you can change policy variables as per your convenience
policy:
  owner:
    age: ">10m" # means owner's age must be greater than "10 months"
    repos: ">3" 
    followers: ">5" 
  repo:
    age: "> 2 y"  # means repo's age must be greater than "2 years"
    stars: ">50"  
    forks: ">2"
    contributors: ">2" 
  commit:
    last_commit_age: "<20d"  # means owner's age must be less than "20 days"
  release:
    last_release: "< 13 d"  # means owner's age must be less than "13 days"
```
### different ways to specify age
| Keyword | Denotes |
| ------- | --------- |
| d | days |
| m | months |
| y | years |


## Understanding Response
Currently 3 types of respone are supported, you use `-rt` or `-resp` flag to get different kind of response.(byDefault: boolean-based)
| ResponseType|Description| FlagValue|
|-------------|-----------|---------- |
|boolean-based| shows whether the policy is followed by package or not, returns either `true` or `false` | not_required |
|json-based| shows detailed output in json format | json |
|yaml-based| show detailed output in yaml format | yaml or yml |
| score | show score of the repo | score |

```json
# sample json-based response
{
 "policy_response": {
  "owner": {
   "age": "<whether_followed_policy>",
   "repos": "<whether_followed_policy>",
   "followers": "<whether_followed_policy>"
  },
  "repo": {
   "age":"<whether_followed_policy>",
   "stars": "<whether_followed_policy>",
   "forks": "<whether_followed_policy>",
   "contributors":"<whether_followed_policy>"
  },
  "commit": {
   "last_commit_age": "<whether_followed_policy>"
  },
  "release": {
   "last_release": "<whether_followed_policy>"
  }
 }
}
```
## Rate limiting
Use `gh_token` environment-variable with ossinspector to avoid `rate-limting` from `github-api`


```bash
  # example
  gh_token=<token_here> ossinspector -p policy.yml -r facebook/react -rt score

```
