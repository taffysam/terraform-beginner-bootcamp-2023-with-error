## Terrahome AWS

### How the module works


Its imponatnt to know the directory ex[ects the following
1. index.html
2. error.html
3. assests

```tf
module "terrahome_aws" {
source = "./modules/terrahome_aws"
user_uuid = var.teacherseat_user_uuid
index_html_file_path = var.index_html_file_path
error_html_file_path = var.error_html_file_path
context_version  = var.content_version
}
```


Note that all top level files in assets will be copied but not any subdirectories