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

import os
import shutil

from os.path import exists as file_exists


class FileSystem:
    """FileSystem Class"""

    def read_file(self, file_path):
        """
        Read file content

        Args:
            file_path: The file path

        Returns:
            The file content
        """

        f = open(file_path, "r")

        return f.read()

    def create_dirs(self, path, mode=0o777):
        """
        Creates a dirs

        Args:
            path: The path
            mode: The mode

        Returns:
            Whether dirs got created or not
        """
        os.makedirs(path, mode)

        return os.path.exists(path)

    def file_exists(self, path):
        """
        Check if file exists

        Args:
            path: the file path

        Returns:
            Whether file exists or not
        """
        return file_exists(path)

    def list_sub_dirs(self, path):
        """
        Get a list of directories

        Args:
            path: the full path

        Returns:
            a list of sub directories
        """
        dirs = []

        for file in os.listdir(path):
            d = os.path.join(path, file)
            if os.path.isdir(d):
                dirs.append(d)

        return dirs

    def change_permission(self, path, mode=0o777):
        """
        Changes file permission

        Args:
            path: The path
            mode: The mode
        """
        os.chmod(path, mode)

    def open_file(self, file_path):
        """
        Open a file for read

        Args:
            file_path: The file path

        Returns:
            the opened file handler
        """

        return open(file_path, "r")

    def write_file(self, file_path, content):
        """
        Write content to a file

        Args:
            file_path: The file path
            content: The file content
        """

        f = open(file_path, "w")
        f.write(content)
        f.close()

    def delete_directory(self, directory):
        """
        Deletes a directory and its content

        Args:
            directory: The directory to delete
        """
        shutil.rmtree(directory)

    def delete_file(self, file_path):
        """
        Delete a file

        Args:
            file_path: The file path
        """
        os.remove(file_path)
