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
import com.cisco.msx.platform.model.IncidentConfig;
import com.cisco.msx.platform.model.IncidentConfigPatch;
import com.cisco.msx.platform.model.IncidentConfigUpdate;
import com.cisco.msx.platform.model.ServiceNowConfiguration;
import com.cisco.msx.platform.model.ServiceNowConfigurationCreate;
import com.cisco.msx.platform.model.ServiceNowConfigurationsPage;
import java.util.UUID;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for IncidentConfigurationsApi
 */
@Ignore
public class IncidentConfigurationsApiTest {

    private final IncidentConfigurationsApi api = new IncidentConfigurationsApi();

    
    /**
     * Creates a ServiceNow configuration.
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createServiceNowConfigurationTest() throws ApiException {
        ServiceNowConfigurationCreate serviceNowConfigurationCreate = null;
        ServiceNowConfiguration response = api.createServiceNowConfiguration(serviceNowConfigurationCreate);

        // TODO: test validations
    }
    
    /**
     * Deletes a ServiceNow configuration.
     *
     * Delete service now configuration, only if no associated entities exists.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteServiceNowConfigurationTest() throws ApiException {
        UUID id = null;
        api.deleteServiceNowConfiguration(id);

        // TODO: test validations
    }
    
    /**
     * API to get configuration of encloud service.
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getConfigurationTest() throws ApiException {
        IncidentConfig response = api.getConfiguration();

        // TODO: test validations
    }
    
    /**
     * Returns a ServiceNow configuration.
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getServiceNowConfigurationTest() throws ApiException {
        UUID id = null;
        ServiceNowConfiguration response = api.getServiceNowConfiguration(id);

        // TODO: test validations
    }
    
    /**
     * Returns a ServiceNow configurations.
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getServiceNowConfigurationsPageTest() throws ApiException {
        Integer page = null;
        Integer pageSize = null;
        UUID tenantId = null;
        String domain = null;
        ServiceNowConfigurationsPage response = api.getServiceNowConfigurationsPage(page, pageSize, tenantId, domain);

        // TODO: test validations
    }
    
    /**
     * API to patch configure encloud service
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void patchConfigurationTest() throws ApiException {
        IncidentConfigPatch incidentConfigPatch = null;
        IncidentConfig response = api.patchConfiguration(incidentConfigPatch);

        // TODO: test validations
    }
    
    /**
     * API to configure encloud service.
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateConfigurationTest() throws ApiException {
        IncidentConfigUpdate incidentConfigUpdate = null;
        IncidentConfig response = api.updateConfiguration(incidentConfigUpdate);

        // TODO: test validations
    }
    
}