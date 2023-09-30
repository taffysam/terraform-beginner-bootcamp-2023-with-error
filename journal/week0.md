Terraform Beginner Boot Camp Week 0

Table of Contents
  * [**Installing Terraform CLI**](#--installing-terraform-cli--)
  * [**Refactoring AWS CLI**](#--refactoring-aws-cli--)
- [Terraform Beginner Boot](#terraform-beginner-boot)
  * [**Semantic Versioning 2.0.0**](#--semantic-versioning-200--)
  * [**Installing Terraform CLI**](#--installing-terraform-cli---1)
  * [Working With Environment Variables](#working-with-environment-variables)
  * [**Terraform Basics**](#--terraform-basics--)
  * [**Terraform Module**](#--terraform-module--)
  * [**Terraform Console**](#--terraform-console--)
  * [**Terraform Key Providers**](#--terraform-key-providers--)
      - [**Key Terraform commands**](#--key-terraform-commands--)
  * [**Terraform Cloud login issues**](#--terraform-cloud-login-issues--)
  * [*Terraform Cloud Login bash Script**](#-terraform-cloud-login-bash-script--)
- [Check if TERRAFORM_CLOUD_TOKEN is set](#check-if-terraform-cloud-token-is-set)
- [Define the JSON structure](#define-the-json-structure)
- [Check if TERRAFORM_CLOUD_TOKEN is set](#check-if-terraform-cloud-token-is-set-1)
- [Specify the target directory and file path](#specify-the-target-directory-and-file-path)
- [Ensure the target directory exists, create it if not](#ensure-the-target-directory-exists--create-it-if-not)
- [Define the JSON structure](#define-the-json-structure-1)
- [Create credentials.tfrc.json](#create-credentialstfrcjson)

<small><i><a href='http://ecotrust-canada.github.io/markdown-toc/'>Table of contents generated with markdown-toc</a></i></small>



**Semantic Versioning 2.0.0**
**Summary**
Given a version number **MAJOR.MINOR.PATCH**, increment the:

- **MAJOR** version when you make incompatible API changes
- **MINOR** version when you add functionality in a backward compatible manner
- **PATCH** version when you make backward compatible bug fixes
Additional labels for pre-release and build metadata are available as extensions to the MAJOR.MINOR.PATCH format.


## **Installing Terraform CLI**
- We have added a new file install_terraform_cli. This is a bash script that will install Terraform CLI on gitpod.
- To have the file run we will change file ownership and as well as make the file executable

- We made use of the below commands 

1. **Changing file ownership** [File Ownership](https://en.wikipedia.org/wiki/Ownership)


2. **Making File Executable** [File-system_permissions](https://en.wikipedia.org/wiki/File-system_permissions)


Also note that when running the bash script add the below line at the top so that you dont have to go via source each tme ypu want to run the file.

**#!/usr/bin/env bash**




## **Refactoring AWS CLI**

- We will refactor AWS CLI installation script into a bash script.
- Provide env variales for AWS CLI requirements
- Set env using gp env

- We made use of the below commands 
# Terraform Beginner Boot

## **Semantic Versioning 2.0.0**
**Summary**
Given a version number **MAJOR.MINOR.PATCH**, increment the:

- **MAJOR** version when you make incompatible API changes
- **MINOR** version when you add functionality in a backward compatible manner
- **PATCH** version when you make backward compatible bug fixes
Additional labels for pre-release and build metadata are available as extensions to the MAJOR.MINOR.PATCH format.


## **Installing Terraform CLI**
- We have added a new file install_terraform_cli. This is a bash script that will install Terraform CLI on gitpod.
- To have the file run we will change file ownership and as well as make the file executable

- We made use of the below commands 

1. **Chaging file ownership** [File Ownership](https://en.wikipedia.org/wiki/Ownership)


2. **Making File Executable** [File-system_permissions](https://en.wikipedia.org/wiki/File-system_permissions)


Also note that when running the bash script add the below line at the top so that you dont have to go via source each tme ypu want to run the file.

**#!/usr/bin/env bash**


## Working With Environment Variables ####

- we can filter variable using env | grep AWS _VAR

**Setting and Unsetting Variables** 
- We can also set and unset variables in git using **export $<>**, *echo &<>* and **unset <>**. Make note that we unsetting the variable we do not use the $.
- Also a variable can be set as below

```
Hello =$Hello
## Printing Variables
echo &AWS_var
##Scoping of Environment Variables
When you open new bash code terminal they will not recognise env variables from the last terminal. So thesewill be set as global or persiting env variables
- When persit env variable in secrect storage.
gp env Hello ="world"

##Environment Variables can also be set in vargitpod.
- NB: These are not supposed to be non sensentive variables


## **AWS CLI INSTALLATION**

AWS CLI is installed for this project and we need via the ./bin/install_aws_cli
[AWS CLI Installation](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)

- Before we start any AWS CLI configuration we can check if any configuration has been done via the below:

```sh
 aws sts get-caller-identity

```

 or 
```sh
aws configure

```
Next step is get AWS CLI envs


## **Terraform Basics**
- get providers, modules and others
- Providers are way you directly interract with an API to make it powered by terraform
- a mapping of APIs to be utilised within terraform.
- direct interfaces to API

## **Terraform Module**
- a collection of terraform files
- used to give a way to utilise the API
- templates 

[These are available at](https://registry.terraform.io/)


## **Terraform Console**
- we can see all the commands by typing 'terraform'




## **Terraform Key Providers**

Example:
Random
- [Created random items](https://registry.terraform.io/providers/hashicorp/random/latest)

#### **Key Terraform commands**

  - **terraform init**          Prepare your working directory for other commands
  - **terraform validate**      Check whether the configuration is valid
  - **terraform plan**          Show changes required by the current configuration
  - **terraform apply**         Create or update infrastructure
  - **terraform destroy**       Destroy previously-created infrastructure
  - **terraform state** 
  - **terraform output** 
  - **terraform apply --auto-approve**  Automatically approve an apply
  
  
  
 **Terraform Lock Files**

  Contains locked versioning for providers and modules to be used for the project.
  Should be committed to your version control system eg GitHub
  
  
  'terraform.tffiles'
   Contains information about your current state of your infrastructure
   Should not be commited to your version control system 
   Can contain sensentive information
   
  'terraform.tffiles.backup'
  the previous state file state


**Terraform Directory** 
  .terraform : contains binaries of terraform

resource "random_string" "random" {
  length           = 16
  special          = true
  override_special = "/@£$"
}


Remote Backend and intergration of Terraform Cloud and state file


**Old Way**

terraform {
  backend "remote" {
    hostname = "app.terraform.io"
    organization = "Tafadzwa"

    workspaces {
      name = "terra-house-1"
    }
  }


[New way](https://developer.hashicorp.com/terraform/language/settings/backends/remote)

cloud {
    organization = "Tafadzwa"

    workspaces {
      name = "terra-house-1"
    }
  }
  
  
Once above code has been added 
  
  CLI-driven runs
Ensure you are properly authenticated into Terraform Cloud by running terraform login on the command line or by using a credentials block.
Add a code block to your Terraform configuration files to set up the cloud integration . 
You can add this configuration block to any .tf file in the directory where you run Terraform.

Run terraform init to initialize the workspace.
Run terraform apply to start the first run for this workspace.

To create a manual token use the link below

https://app.terraform.io/app/settings/tokens?source=terraform-login

## **Terraform Cloud login issues**

When running terraform-login

When faced with issues you can add the token to the below file /home/gitpod/.terraform.d/credentials.tfrc.json

touch /home/gitpod/.terraform.d/credentials.tfrc.json

"credentials": {
    "app.terraform.io": {
      "token": 'YyRhq3UCnudUYw.$%%%@hblabdSLWITgggggdeeewLPzsVgsRffdfl4yN2cCDNILz2Ik5xhhnillCHa99cBKyLbExapmle'
    }
}

Above key is generated from To create a manual token use the link below

https://app.terraform.io/app/settings/tokens?source=terraform-login



## *Terraform Cloud Login bash Script**

ChatGPT created the below bash script 


#!/bin/bash

# Check if TERRAFORM_CLOUD_TOKEN is set
if [ -z "$TERRAFORM_CLOUD_TOKEN" ]; then
  echo "Error: TERRAFORM_CLOUD_TOKEN environment variable is not set."
  exit 1
fi

# Define the JSON structure
```sh
json_data='{
  "credentials": {
    "app.terraform.io": {
      "token": "'"$TERRAFORM_CLOUD_TOKEN"'"
    }
  }
}'

# Create credentials.tfrc.json
echo "$json_data" > credentials.tfrc.json

echo "credentials.tfrc.json file has been generated."

Save this script to a file, e.g., generate_credentials.sh, and make it executable

chmod +x generate_credentials.sh

```

Now, when you run the script, it will generate a credentials.tfrc.json file with the specified JSON structure using the TERRAFORM_CLOUD_TOKEN environment variable:

```sh
./generate_credentials.sh

```

An edit to the above script was done to do the below:

This script will:

    Check if the TERRAFORM_CLOUD_TOKEN environment variable is set.
    Specify the target directory and file path, creating the directory if it doesn't exist.
    Define the JSON structure with the token.
    Create the credentials.tfrc.json file in the specified location.
	


New Bash Script:#!/bin/bash

# Check if TERRAFORM_CLOUD_TOKEN is set
if [ -z "$TERRAFORM_CLOUD_TOKEN" ]; then
  echo "Error: TERRAFORM_CLOUD_TOKEN environment variable is not set."
  exit 1
fi

# Specify the target directory and file path
target_directory="/home/gitpod/.terraform.d"
target_file="$target_directory/credentials.tfrc.json"

# Ensure the target directory exists, create it if not
mkdir -p "$target_directory"

# Define the JSON structure
json_data='{
  "credentials": {
    "app.terraform.io": {
      "token": "'"$TERRAFORM_CLOUD_TOKEN"'"
    }
  }
}'

# Create credentials.tfrc.json
echo "$json_data" > "$target_file"

echo "credentials.tfrc.json file has been generated in $target_file."
	



##Setting Alias

```sh
open ~/.bash_profile 

alias tf ="terraform"
```


##To reload the above changes to be effected
source ~/.bash_profile 

gp env TERRAFORM_CLOUD_TOKEN='EXAMPLSTOkenLEtsTRYadneeeek5xhCHa99cBKyLb1P9zm4mkQ'



This script does the following:

    - Defines the alias tf='terraform'.

    - Determines the user's home directory.

    - Checks if the .bash_profile file exists.

    - Adds the alias definition to the .bash_profile file if it doesn't already exist.

    - Sources the .bash_profile file to apply the changes immediately.

You can save this script to a file, e.g., add_tf_alias.sh, make it executable with chmod +x add_tf_alias.sh, and then run it.

 After running the script, the tf alias will be available in your terminal sessions.
 
  
 The Bash Script :

```sh
 
#!/bin/bash

# Define the alias
alias tf='terraform'

# Get the user's home directory
user_home_dir="$HOME"

# Check if the .bash_profile file exists
bash_profile="$user_home_dir/.bash_profile"

# Add the alias definition to .bash_profile if not already present
if ! grep -q "alias tf='terraform'" "$bash_profile"; then
  echo "alias tf='terraform'" >> "$bash_profile"
  echo "Alias 'tf' added to $bash_profile."
else
  echo "Alias 'tf' already exists in $bash_profile."
fi

# Source .bash_profile to apply the changes immediately
source "$bash_profile"

echo "Alias 'tf' has been added. You can now use 'tf' as an alias for 'terraform'."
```
