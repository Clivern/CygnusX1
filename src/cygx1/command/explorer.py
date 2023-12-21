# MIT License
#
# Copyright (c) 2022 Clivern
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

import os
import click
from datetime import datetime

from cygx1.module.facts import Facts
from cygx1.module.nasa import Nasa
from cygx1.module.logger import Logger
from cygx1.module.checksum import Checksum
from cygx1.module.archive import Archive
from cygx1.module.file_system import FileSystem


class Launch:
    """Launch Class"""

    def __init__(self, file, platform):
        self.file = file
        self.nasa = Nasa()
        self.platform = platform
        self.file_system = FileSystem()
        self.checksum = Checksum()
        self.facts = Facts()
        self.archive = Archive("archive.json")
        self.logger = Logger().get_logger(__name__)

    def run(self):
        """Run explorer"""
        now = datetime.now()
        formatted_date_time = now.strftime("%Y-%m-%d %H:%M:%S")

        fact = self.facts.random_line('facts.txt').strip()

        if self.platform == "nasa":
            result = self.nasa.fetch(os.getenv('NASA_API_KEY'))
            data = f"""<p align='center'>
  <img src='{result['url']}' width='60%' />
    <h3 align="center">Cygnus X-1</h3>
    <p align="center">{fact}</p>
</p>
<br/>

Explanation
--
{result['explanation']}


*Last updated at {formatted_date_time}*
"""
            data = data.replace("   Explore Your Universe: Random APOD Generator", "")
            self.file_system.write_file(self.file, data)
            checksum = self.checksum.measure(result['explanation'])
            self.archive.append(checksum, result['url'], result['explanation'])

        click.echo("Explorer finished successfully!")
