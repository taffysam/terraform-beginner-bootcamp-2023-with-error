## Terraform Beginner Bootcamp 2023

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



## Fix missing Resources with terraform import

If someone goes and delete or modifies cloud resources manually through clickOps. If we run the terraform plan again it will attempt to put
infrastructure back into original state. 

If you loose you state you have to tear down all infrastructure manually, you will need to run terraform import. 



[Terraform Import](https://developer.hashicorp.com/terraform/language/import)


## Passing Input Variables
- We can pass input variable to out modul.

## Module source
Using the source we can import module from various places ge:
- locally
- GitHub
- Terraform Registry

```terraform
module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
  user_uuid = var.user_uuid
  bucket_name = var.bucket_name
}


```


## Terraform Import




## Considerations when using ChatGPT to write terraform
ChatGPT may not be trained on the latest version of terraform. Providers to change from time to time 


[Referencing Files](https://developer.hashicorp.com/terraform/language/expressions/references)

 
[Module Sources](https://developer.hashicorp.com/terraform/language/modules/sources) 

In terraform there is a special varaible called `path` for the variable path that allows us to reference local path.

## The following values are available:

    - path.module is the filesystem path of the module where the expression is placed. We do not recommend using path.module 
	in write operations because it can produce different behavior depending on whether you use remote or local module sources. 
	Multiple invocations of local modules use the same source directory, overwriting the data in path.module during each call. 
	This can lead to race conditions and unexpected results.
    - path.root is the filesystem path of the root module of the configuration.
    - path.cwd is the filesystem path of the original working directory from where you ran Terraform before applying any -chdir argument. 
	This path is an absolute path that includes details about the filesystem structure. It is also useful in some advanced cases where 
	Terraform is run from a directory other than the root module directory. We recommend using path.root or path.module over path.cwd where possible.
    - terraform.workspace is the name of the currently selected workspace.


## The main kinds of named values available in Terraform are:

    - Resources
    - Input variables
    - Local values
    - Child module outputs
    - Data sources
    - Filesystem and workspace info
    - Block-local values



# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_object
resource "aws_s3_object" "object" {
  bucket = "87e8fc20-5f21-4b38-872b-ab8adfb49ed5"
  key    = "index.html"
  
  source = var.index_html_file_path
    
  source = "${path.root}/public.index.html"
  
  source = "c:/Sofwares/Terraform Projects/terraform-beginner-bootcamp-2023-1/modules/terrahouse_aws/public/index.html"

    # The filemd5() function is available in Terraform 0.11.12 and later
  # For Terraform 0.11.11 and earlier, use the md5() function and the file() function:
  # etag = "${md5(file("path/to/file"))}"
  # etag = filemd5("path/to/file")



  ${path.root}/public/index.html
  
 etag = filemd5("${path.root}/public/index.html)
 
 
  soource= var.index.html.file_path

}


## Terraform Data Sources

[Data Source](https://developer.hashicorp.com/terraform/language/data-sources)


[AWS Data caller ID](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity)



## Terraform Locals

Locals allow us to destry locl variables. It canbe very usefulwhen we need to transform data into another format and

```tf
locals {
  s3_origin_id = "MyS3Origin"
}
```


## Provisioners
[Provisioners](https://developer.hashicorp.com/terraform/language/resources/provisioners/file)
- Allow you to execute local commnands on compute instances eg and AWS CLI Command
- Not reccommended by Harshicorp instead use Ansible with better functionality

### Local-execs
- This will execute commands on the commands running on machine on which you target. You will need to provide credentials such as ssh to
get into the machine.


### Remote-execs


### File Provisioners

-The file provisioner copies files or directories from the machine running Terraform to the newly created resource. The file provisioner supports both ssh and winrm type

`Example`

```tf
resource "aws_instance" "web" {
  # ...

  # Copies the myapp.conf file to /etc/myapp.conf
  provisioner "file" {
    source      = "conf/myapp.conf"
    destination = "/etc/myapp.conf"
  }

  # Copies the string in content into /tmp/file.log
  provisioner "file" {
    content     = "ami used: ${self.ami}"
    destination = "/tmp/file.log"
  }

  # Copies the configs.d folder to /etc/configs.d
  provisioner "file" {
    source      = "conf/configs.d"
    destination = "/etc"
  }

  # Copies all files and folders in apps/app1 to D:/IIS/webapp1
  provisioner "file" {
    source      = "apps/app1/"
    destination = "D:/IIS/webapp1"
  }
}


```

## Invalidating Cloud Front 

Use the below code

resource "aws_lambda_function" "cloudfront_invalidation" {
  filename         = "cloudfront_invalidation.zip"
  source_code_hash = filebase64sha256("cloudfront_invalidation.zip")
  role             = aws_iam_role.lambda_exec.arn

  handler = "index.handler"
  runtime = "nodejs14.x"
}

data "aws_iam_policy_document" "lambda" {
  source_json = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "cloudfront:CreateInvalidation",
        Effect = "Allow",
        Resource = aws_cloudfront_distribution.my_distribution.arn,
      },
    ],
  })
}

resource "aws_iam_role" "lambda_exec" {
  name = "lambda_exec_role"
}

resource "aws_iam_policy" "lambda" {
  name        = "lambda_policy"
  description = "Policy for CloudFront Invalidation Lambda"
  policy      = data.aws_iam_policy_document.lambda.json
}

resource "aws_iam_role_policy_attachment" "lambda" {
  policy_arn = aws_iam_policy.lambda.arn
  role       = aws_iam_role.lambda_exec.name
}

resource "aws_lambda_permission" "cloudfront_invalidation" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.cloudfront_invalidation.arn
  principal     = "cloudfront.amazonaws.com"
  source_arn    = aws_cloudfront_distribution.my_distribution.arn
}

resource "null_resource" "invalidate" {
  triggers = {
    function_version = aws_lambda_function.cloudfront_invalidation.qualified_arn
  }

  provisioner "local-exec" {
    command = "aws lambda update-function-code --function-name ${aws_lambda_function.cloudfront_invalidation.function_name} --zip-file fileb://${path.module}/cloudfront_invalidation.zip"
  }
}


## Heredocs
[Heredocs](https://developer.hashicorp.com/terraform/language/expressions/strings)

Can be used when doing multi-line commands
`Syntax`
<<EOT
hello
world
EOT














## Security on Terraform
### Managing Secrets

### Key files to be protected
- State File
- Terraform Variable files
- Terraform Providers
- Terraform instances

`Declare variables and sensetive variables locally`
`Access to Sensetive Data on Terraform can be managed by Specific Users`
`Can be Cloud native`
- Aws Secret Manager
- Azure Key Vault
- Harshicorp Vault
- Encryption 

Variables can be Terraform or Environmental



## For Each

 - We use the for_each attribute to loop through each file in the local.asset_files list.
 - The bucket attribute references the S3 bucket you created.
 - The key attribute is the name of the object in the S3 bucket, which is the same as the filename.
 - The source attribute is the local file path to upload.
 - The acl attribute sets the object's access control list to "private." You can change it to your desired ACL.
 
 ### Sample Code
 
 ```tf
 resource "aws_s3_bucket" "example" {
  bucket = "your-s3-bucket-name"
  acl    = "private"
}

 
 ```
 
 
 
 ```tf
 variable "assets_folder_path" {
  description = "Path to the 'assets' folder"
  default     = "./assets"
}

locals {
  asset_files = fileset(var.assets_folder_path, "**/*")
}
 
 
 ```
 
 
 
 ```tf
 
 resource "aws_s3_bucket_object" "asset_objects" {
  for_each = { for file in local.asset_files : file => file }

  bucket = aws_s3_bucket.example.id
  key    = each.value
  source = "${var.assets_folder_path}/${each.value}"
  acl    = "private"
}

 
 ```
 
 
 ## Collection types
 
 List:
 
 
 Map:
 
 
 Set:
 
 
 ## CRUD

 CRUD stands for Create, Read, Update and Delete

 [CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete)




 