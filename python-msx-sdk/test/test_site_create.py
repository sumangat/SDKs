"""
    MSX SDK

    MSX SDK client.  # noqa: E501

    The version of the OpenAPI document: 1.0.9
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import python_msx_sdk
from python_msx_sdk.model.site_address import SiteAddress
from python_msx_sdk.model.site_contact import SiteContact
from python_msx_sdk.model.site_create_all_of import SiteCreateAllOf
from python_msx_sdk.model.site_location import SiteLocation
from python_msx_sdk.model.site_update import SiteUpdate
globals()['SiteAddress'] = SiteAddress
globals()['SiteContact'] = SiteContact
globals()['SiteCreateAllOf'] = SiteCreateAllOf
globals()['SiteLocation'] = SiteLocation
globals()['SiteUpdate'] = SiteUpdate
from python_msx_sdk.model.site_create import SiteCreate


class TestSiteCreate(unittest.TestCase):
    """SiteCreate unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testSiteCreate(self):
        """Test SiteCreate"""
        # FIXME: construct object with mandatory attributes with example values
        # model = SiteCreate()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
