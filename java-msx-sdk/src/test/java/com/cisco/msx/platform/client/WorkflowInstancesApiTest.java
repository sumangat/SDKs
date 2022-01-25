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
import java.time.LocalDate;
import com.cisco.msx.platform.model.WorkflowAction;
import com.cisco.msx.platform.model.WorkflowInstance;
import com.cisco.msx.platform.model.WorkflowInstanceDeleteResponse;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for WorkflowInstancesApi
 */
@Ignore
public class WorkflowInstancesApiTest {

    private final WorkflowInstancesApi api = new WorkflowInstancesApi();

    
    /**
     * Cancels a workflow instance.
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void cancelWorkflowInstanceTest() throws ApiException {
        String id = null;
        WorkflowInstance response = api.cancelWorkflowInstance(id);

        // TODO: test validations
    }
    
    /**
     * Deletes a workflow instance.
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteWorkflowInstanceTest() throws ApiException {
        String id = null;
        WorkflowInstanceDeleteResponse response = api.deleteWorkflowInstance(id);

        // TODO: test validations
    }
    
    /**
     * Returns a workflow instance.
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getWorkflowInstanceTest() throws ApiException {
        String id = null;
        WorkflowInstance response = api.getWorkflowInstance(id);

        // TODO: test validations
    }
    
    /**
     * Returns a workflow instance action.
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getWorkflowInstanceActionTest() throws ApiException {
        String id = null;
        String actionId = null;
        WorkflowAction response = api.getWorkflowInstanceAction(id, actionId);

        // TODO: test validations
    }
    
    /**
     * Returns a list of workflow instances.
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getWorkflowInstancesListTest() throws ApiException {
        String id = null;
        Integer page = null;
        Integer pageSize = null;
        LocalDate dateFrom = null;
        LocalDate dateTo = null;
        List<WorkflowInstance> response = api.getWorkflowInstancesList(id, page, pageSize, dateFrom, dateTo);

        // TODO: test validations
    }
    
}