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


class Host:
    """Host Model"""

    def __init__(
        self,
        id,
        name,
        connection,
        ip,
        port,
        user,
        password,
        ssh_private_key,
        tags,
        created_at,
        updated_at,
    ):
        """Class Constructor"""
        self._id = id
        self._name = name
        self._connection = connection
        self._ip = ip
        self._port = port
        self._user = user
        self._password = password
        self._ssh_private_key = ssh_private_key
        self._tags = tags
        self._created_at = created_at
        self._updated_at = updated_at

    @property
    def id(self):
        """Host ID"""
        return self._id

    @property
    def name(self):
        """Host Name"""
        return self._name

    @property
    def connection(self):
        """Host Connection"""
        return self._connection

    @property
    def ip(self):
        """Host IP"""
        return self._ip

    @property
    def port(self):
        """Host Port"""
        return self._port

    @property
    def user(self):
        """Host User"""
        return self._user

    @property
    def password(self):
        """Host Password"""
        return self._password

    @property
    def ssh_private_key(self):
        """Host SSH Key"""
        return self._ssh_private_key

    @property
    def tags(self):
        """Host Tags"""
        return self._tags

    @user.setter
    def user(self, user):
        """Set User"""
        self._user = user

    @password.setter
    def password(self, password):
        """Set Password"""
        self._password = password

    @ssh_private_key.setter
    def ssh_private_key(self, ssh_private_key):
        """Set SSH Private Key"""
        self._ssh_private_key = ssh_private_key

    @property
    def created_at(self):
        """Host Created At"""
        return self._created_at

    @property
    def updated_at(self):
        """Host Updated At"""
        return self._updated_at
