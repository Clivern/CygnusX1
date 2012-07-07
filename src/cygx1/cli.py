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

import click

from cygx1 import __version__
from cygx1.command.explorer import Launch


@click.group(help="🐺 The immensity of the universe is difficult to grasp.")
@click.version_option(version=__version__, help="Show the current version")
def main():
    pass


# Explorer command
@click.group(help="Explorer Commands")
def explorer():
    pass


# Launch sub command
@explorer.command(help="Launch sub Command")
@click.argument("file")
@click.option("-p", "--platform", "platform", type=click.STRING, default="nasa", help="the platform")
def launch(file, platform):
    return Launch(file, platform).run()


# Register Commands
main.add_command(explorer)


if __name__ == "__main__":
    main()
