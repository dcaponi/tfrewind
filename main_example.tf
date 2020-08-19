provider aws {
	alias = "aws"
}

resource aws_instance main {
	provider = aws
	instance_type = "t2.micro"
	ami = "ami-<id>"
	availability_zone = "<region>"
	cpu_core_count = 1
	cpu_threads_per_core = 1
	tenancy = "default"
}

resource aws_security_group main {
	provider = aws
	name = "default"
	revoke_rules_on_delete = false
	vpc_id = "vpc-<id>"
	description = "default VPC security group"
}

resource aws_subnet main {
	provider = aws
	cidr_block = "<x.x.x.x/y>"
	vpc_id = "vpc-<id>"
	availability_zone = "<region>"
}

resource aws_vpc main {
	provider = aws
	cidr_block = "x.x.x.x/y"
	enable_dns_hostnames = true
	enable_dns_support = true
	instance_tenancy = "default"
}

