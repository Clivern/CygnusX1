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

import click

from jaglion.module.logger import Logger
from jaglion.module.output import Output
from jaglion.module.config import Config
from jaglion.module.encrypt import Encrypt
from jaglion.module.database import Database
from jaglion.module.file_system import FileSystem


class Hosts:
    """Hosts Class"""

    def __init__(self):
        self.output = Output()
        self.database = Database()
        self.config = Config()
        self.encrypt = Encrypt()
        self.file_system = FileSystem()
        self.logger = Logger().get_logger(__name__)

    def init(self):
        """Init database and configs"""
        self.configs = self.config.load()
        self.database.connect(self.configs["database"]["path"])
        self.database.migrate()
        return self

    def delete(self, name):
        """Delete a host"""
        self.database.delete_host(name)

        click.echo(f"Host with name {name} got deleted")
