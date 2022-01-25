/*
 * MSX SDK
 * MSX SDK client.
 *
 * The version of the OpenAPI document: 1.0.9
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package com.cisco.msx.platform.client;

import com.cisco.msx.platform.ApiException;
import com.cisco.msx.platform.model.Error;
import com.cisco.msx.platform.model.LicenseSummary;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for LicensingLicensesApi
 */
@Ignore
public class LicensingLicensesApiTest {

    private final LicensingLicensesApi api = new LicensingLicensesApi();

    
    /**
     * Returns a filtered list of licenses.
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getLicensesListTest() throws ApiException {
        String smartAccountId = null;
        String virtualAccountId = null;
        List<LicenseSummary> response = api.getLicensesList(smartAccountId, virtualAccountId);

        // TODO: test validations
    }
    
}