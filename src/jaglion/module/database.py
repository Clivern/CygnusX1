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

import json
import sqlite3

from jaglion.model.host import Host
from jaglion.model.secret import Secret


class Database:
    """Database Class"""

    def connect(self, path):
        """Connect into a database"""
        self.path = path

        self._connection = sqlite3.connect(self.path)

        return self._connection.total_changes

    def migrate(self):
        cursor = self._connection.cursor()

        cursor.execute(
            "CREATE TABLE IF NOT EXISTS host (id TEXT, name TEXT, config TEXT, createdAt TEXT, updatedAt TEXT)"
        )
        cursor.execute(
            "CREATE TABLE IF NOT EXISTS task (id TEXT, name TEXT, payload TEXT, result TEXT, createdAt TEXT, updatedAt TEXT)"
        )
        cursor.execute(
            "CREATE TABLE IF NOT EXISTS secret (id TEXT, name TEXT, value TEXT, createdAt TEXT, updatedAt TEXT)"
        )

        cursor.close()

        self._connection.commit()

    def get_host(self, name):
        """Get a row by host name"""
        cursor = self._connection.cursor()

        rows = cursor.execute(
            "SELECT id, name, config, createdAt, updatedAt FROM host WHERE name = '{}'".format(
                name
            )
        ).fetchall()

        cursor.close()

        if len(rows) > 0:
            for row in rows:
                data = json.loads(row[2])

                host = Host(
                    row[0],
                    row[1],
                    data["connection"],
                    data["ip"],
                    data["port"],
                    data["user"],
                    data["password"],
                    data["ssh_private_key"],
                    data["tags"],
                    row[3],
                    row[4],
                )

                return host
        else:
            return None

    def insert_host(self, host):
        """Insert a new row"""
        cursor = self._connection.cursor()

        result = cursor.execute(
            "INSERT INTO host VALUES ('{}', '{}', '{}', datetime('now'), datetime('now'))".format(
                host.id,
                host.name,
                json.dumps(
                    {
                        "connection": host.connection,
                        "ip": host.ip,
                        "port": host.port,
                        "user": host.user,
                        "password": host.password,
                        "ssh_private_key": host.ssh_private_key,
                        "tags": host.tags,
                    }
                ),
            )
        )

        cursor.close()

        self._connection.commit()

        return result.rowcount

    def list_hosts(self):
        """List all rows"""
        result = []

        cursor = self._connection.cursor()
        rows = cursor.execute(
            "SELECT id, name, config, createdAt, updatedAt FROM host"
        ).fetchall()
        cursor.close()

        for row in rows:
            data = json.loads(row[2])

            host = Host(
                row[0],
                row[1],
                data["connection"],
                data["ip"],
                data["port"],
                data["user"],
                data["password"],
                data["ssh_private_key"],
                data["tags"],
                row[3],
                row[4],
            )

            result.append(host)

        return result

    def delete_host(self, name):
        """Delete a row by host name"""
        cursor = self._connection.cursor()

        cursor.execute("DELETE FROM host WHERE name = ?", (name,))

        cursor.close()

        self._connection.commit()

    def get_secret(self, name):
        """Get a row by secret name"""
        cursor = self._connection.cursor()

        rows = cursor.execute(
            "SELECT id, name, value, createdAt, updatedAt FROM secret WHERE name = '{}'".format(
                name
            )
        ).fetchall()

        cursor.close()

        if len(rows) > 0:
            for row in rows:
                data = json.loads(row[2])

                secret = Secret(
                    row[0],
                    row[1],
                    data["value"],
                    data["tags"],
                    row[3],
                    row[4],
                )

                return secret
        else:
            return None

    def list_secrets(self):
        """List all rows"""
        result = []

        cursor = self._connection.cursor()

        rows = cursor.execute(
            "SELECT id, name, value, createdAt, updatedAt FROM secret"
        ).fetchall()

        cursor.close()

        for row in rows:
            data = json.loads(row[2])

            secret = Secret(
                row[0],
                row[1],
                data["value"],
                data["tags"],
                row[3],
                row[4],
            )

            result.append(secret)

        return result

    def insert_secret(self, secret):
        """Insert a new row"""
        cursor = self._connection.cursor()

        result = cursor.execute(
            "INSERT INTO secret VALUES ('{}', '{}', '{}', datetime('now'), datetime('now'))".format(
                secret.id,
                secret.name,
                json.dumps(
                    {
                        "value": secret.value,
                        "tags": secret.tags,
                    }
                ),
            )
        )

        cursor.close()

        self._connection.commit()

        return result.rowcount

    def delete_secret(self, name):
        """Delete a row by secret name"""
        cursor = self._connection.cursor()

        cursor.execute("DELETE FROM secret WHERE name = ?", (name,))

        cursor.close()

        self._connection.commit()
