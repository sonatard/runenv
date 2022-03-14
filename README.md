# runenv

`runenv` create gcloud run deploy `--set-env-vars=` option and export shell environment from yaml file.

## Motivation

I want to manage Cloud Run environment variables in a yaml file. But there is no way to pass the file to [gcloud run deploy](https://cloud.google.com/sdk/gcloud/reference/run/deploy).
Also, in [gcloud services replace](https://cloud.google.com/run/docs/configuring/environment-variables), it can manage environment variables in a yaml file, but it has to manage other values as well. In particular, we need to specify the revision, which is difficult to use.

- Mar 14, 2022 Update

`gcloud beta deploy` supports `--env-vars-file`. https://cloud.google.com/sdk/gcloud/reference/run/deploy

```yaml
KEY1: value1
KEY2: value2
```


## Install

```
go install github.com/sonatard/runenv@latest
```

## Usage

```yaml
# env.yaml
- name: KEY1
  value: value1
- name: KEY2
  value: value2
```

```
$ runenv env.yaml
KEY1=value1,KEY2=value2

$ runenv -e env.yaml
export KEY1="value1" KEY2="value2"
```

## Example

### gcloud run deploy --set-env-vars

1. Prepare revision config

```
gcloud run services describe [SERVICE] --format=export > service.yaml
```

2. Create env.yaml from service.yaml

```
brew install yq
yq ".spec.template.spec.containers[].env" service.yaml > env.yaml

# or you can create env.yaml manualy.
```

```yaml
# env.yaml
- name: KEY1
  value: value1
- name: KEY2
  value: value2
```

3. Cloud run deploy with environment variables

```
gcloud run deploy [SERVICE] --image=[IMAGE] --set-env-vars="$(runenv env.yaml)"
```

### Export shell environment

1. Create env.yaml

```yaml
# env.yaml
- name: KEY1
  value: value1
- name: KEY2
  value: value2
```

2. Export shell environment

```
$ eval "$(runenv -e env.yaml)"
$ echo $KEY1
value1
```
