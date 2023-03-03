# MIT License
#
# Copyright (c) 2023 Clivern
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

import uuid
import click
import logging, json, sys

from jaglion import __version__
from jaglion.model.host import Host
from jaglion.model.secret import Secret
from jaglion.command.hosts import Hosts
from jaglion.command.secret import Secrets
from jaglion.command.configs import Configs


@click.group(help="🐺 Modern Command Line Tool for Apache Kafka")
@click.version_option(version=__version__, help="Show the current version")
def main():
    pass


# Hosts command
@click.group(help="Manage hosts")
def host():
    pass


# Delete host sub command
@host.command(help="Delete a host")
@click.argument("name")
def delete(name):
    return Hosts().init().delete(name)


# Manage configs command
@click.group(help="Manage configs")
def config():
    pass


# Init configs sub command
@config.command(help="Init configurations")
def init():
    return Configs().init()


# Edit configs sub command
@config.command(help="Edit configurations")
def edit():
    return Configs().edit()


# Show configs sub command
@config.command(help="Show configurations")
def dump():
    return Configs().dump()


# Secrets command
@click.group(help="Manage secrets")
def secret():
    pass


# Add secrets sub command
@secret.command(help="Add a secret")
@click.argument("name")
@click.argument("value")
@click.option("-t", "--tags", "tags", type=click.STRING, default="", help="Secret tags")
@click.option("-f", "--force", "force", is_flag=True, default=False, help="Force add")
def add(name, value, tags, force):
    secret = Secret(
        str(uuid.uuid4()),
        name,
        value,
        tags.split(",") if tags != "" else [],
        None,
        None,
    )

    return Secrets().init().add(secret, force)


# List secrets sub command
@secret.command(help="List all secrets")
@click.option("-t", "--tag", "tag", type=click.STRING, default="", help="Secret tag")
@click.option(
    "-o", "--output", "output", type=click.STRING, default="", help="Output format"
)
def list(tag, output):
    return Secrets().init().list(tag, output)


# Get secret sub command
@secret.command(help="Get a secret")
@click.argument("name")
@click.option(
    "-o", "--output", "output", type=click.STRING, default="", help="Output format"
)
def get(name, output):
    return Secrets().init().get(name, output)


# Delete secret sub command
@secret.command(help="Delete a secret")
@click.argument("name")
def delete(name):
    return Secrets().init().delete(name)


# Register Commands
main.add_command(host)
main.add_command(config)
main.add_command(secret)


if __name__ == "__main__":
    main()
