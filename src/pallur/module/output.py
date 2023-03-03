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
from re import sub
from prettytable import PrettyTable


class Output:
    """Output Class"""

    JSON = "JSON"
    DEFAULT = "DEFAULT"

    def _table(self, data):
        """Output data as Table"""
        headers = []
        rows = []

        for item in data:
            headers = item.keys()
            rows.append(item.values())

        x = PrettyTable()
        x.field_names = headers
        x.add_rows(rows)

        return x

    def _json(self, data):
        """Output data as JSON"""
        new = []

        for item in data:
            new.append({self.camel_case(k): v for k, v in item.items()})

        return json.dumps(new)

    def render(self, data, typ):
        """Render Data to the Console"""
        if typ == Output.JSON:
            return self._json(data)

        return self._table(data)

    def camel_case(self, value):
        """Change string into camel case"""
        value = sub(r"(_|-)+", " ", value).title().replace(" ", "")

        return "".join([value[0].lower(), value[1:]])
