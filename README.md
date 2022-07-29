# ossinspector
A policy based open source package inspector 


## Build

```bash

  git clone https://github.com/h0x0er/ossinspector.git
  cd ossinspector
  docker build -t inspector .
 
```

## Usage


```bash
  docker run inspector -h
```
![image](https://user-images.githubusercontent.com/84621253/180490085-b1c9a5b3-1601-45c8-ba97-6393cfebc042.png)

```bash
  docker run inspector -policy main/policy.yml -repo facebook/react -rt json
```
![image](https://user-images.githubusercontent.com/84621253/180490735-3cb93a9b-1c9b-407e-a9b3-83d9c722120e.png)


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
