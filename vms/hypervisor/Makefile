all: qemu aws

.PHONY: qemu
qemu:
	packer build -only=qemu -var-file=template.vars template.json

.PHONY: aws
aws:
	packer build -only=amazon-ebs -var-file=template.vars -var aws_access_key=${AWS_ACCESS_KEY} -var aws_secret_key=${AWS_SECRET_KEY} template.json
