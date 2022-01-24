"""
    MSX SDK

    MSX SDK client.  # noqa: E501

    The version of the OpenAPI document: 1.0.9
    Generated by: https://openapi-generator.tech
"""


import unittest

import python_msx_sdk
from python_msx_sdk.api.templates_api import TemplatesApi  # noqa: E501


class TestTemplatesApi(unittest.TestCase):
    """TemplatesApi unit test stubs"""

    def setUp(self):
        self.api = TemplatesApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_delete_template(self):
        """Test case for delete_template

        Deletes a template.  # noqa: E501
        """
        pass

    def test_get_template(self):
        """Test case for get_template

        Returns a template by id.  # noqa: E501
        """
        pass

    def test_get_template_history(self):
        """Test case for get_template_history

        Returns a template history by id.  # noqa: E501
        """
        pass

    def test_get_templates_page(self):
        """Test case for get_templates_page

        Returns a page of templates.  # noqa: E501
        """
        pass

    def test_import_template(self):
        """Test case for import_template

        Imports a template.  # noqa: E501
        """
        pass

    def test_update_template_status(self):
        """Test case for update_template_status

        Updates a template status.  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
