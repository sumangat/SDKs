/**
 * MSX SDK
 * MSX SDK client.
 *
 * The version of the OpenAPI document: 1.0.9
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from "../ApiClient";
import Error from '../model/Error';
import IncidentConfig from '../model/IncidentConfig';
import IncidentConfigPatch from '../model/IncidentConfigPatch';
import IncidentConfigUpdate from '../model/IncidentConfigUpdate';
import ServiceNowConfiguration from '../model/ServiceNowConfiguration';
import ServiceNowConfigurationCreate from '../model/ServiceNowConfigurationCreate';
import ServiceNowConfigurationsPage from '../model/ServiceNowConfigurationsPage';

/**
* IncidentConfigurations service.
* @module api/IncidentConfigurationsApi
* @version 1.0.9
*/
export default class IncidentConfigurationsApi {

    /**
    * Constructs a new IncidentConfigurationsApi. 
    * @alias module:api/IncidentConfigurationsApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }



    /**
     * Creates a ServiceNow configuration.
     * @param {module:model/ServiceNowConfigurationCreate} serviceNowConfigurationCreate 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/ServiceNowConfiguration} and HTTP response
     */
    createServiceNowConfigurationWithHttpInfo(serviceNowConfigurationCreate) {
      let postBody = serviceNowConfigurationCreate;
      // verify the required parameter 'serviceNowConfigurationCreate' is set
      if (serviceNowConfigurationCreate === undefined || serviceNowConfigurationCreate === null) {
        throw new Error("Missing the required parameter 'serviceNowConfigurationCreate' when calling createServiceNowConfiguration");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = ServiceNowConfiguration;
      return this.apiClient.callApi(
        '/incident/api/v8/configurations', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Creates a ServiceNow configuration.
     * @param {module:model/ServiceNowConfigurationCreate} serviceNowConfigurationCreate 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/ServiceNowConfiguration}
     */
    createServiceNowConfiguration(serviceNowConfigurationCreate) {
      return this.createServiceNowConfigurationWithHttpInfo(serviceNowConfigurationCreate)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Deletes a ServiceNow configuration.
     * Delete service now configuration, only if no associated entities exists.
     * @param {String} id 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing HTTP response
     */
    deleteServiceNowConfigurationWithHttpInfo(id) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling deleteServiceNowConfiguration");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = null;
      return this.apiClient.callApi(
        '/incident/api/v8/configurations/{id}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Deletes a ServiceNow configuration.
     * Delete service now configuration, only if no associated entities exists.
     * @param {String} id 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}
     */
    deleteServiceNowConfiguration(id) {
      return this.deleteServiceNowConfigurationWithHttpInfo(id)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * API to get configuration of encloud service.
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/IncidentConfig} and HTTP response
     */
    getConfigurationWithHttpInfo() {
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = IncidentConfig;
      return this.apiClient.callApi(
        '/incident/api/v1/config', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * API to get configuration of encloud service.
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/IncidentConfig}
     */
    getConfiguration() {
      return this.getConfigurationWithHttpInfo()
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Returns a ServiceNow configuration.
     * @param {String} id 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/ServiceNowConfiguration} and HTTP response
     */
    getServiceNowConfigurationWithHttpInfo(id) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getServiceNowConfiguration");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = ServiceNowConfiguration;
      return this.apiClient.callApi(
        '/incident/api/v8/configurations/{id}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Returns a ServiceNow configuration.
     * @param {String} id 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/ServiceNowConfiguration}
     */
    getServiceNowConfiguration(id) {
      return this.getServiceNowConfigurationWithHttpInfo(id)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Returns a ServiceNow configurations.
     * @param {Number} page 
     * @param {Number} pageSize 
     * @param {Object} opts Optional parameters
     * @param {String} opts.tenantId 
     * @param {String} opts.domain 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/ServiceNowConfigurationsPage} and HTTP response
     */
    getServiceNowConfigurationsPageWithHttpInfo(page, pageSize, opts) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'page' is set
      if (page === undefined || page === null) {
        throw new Error("Missing the required parameter 'page' when calling getServiceNowConfigurationsPage");
      }
      // verify the required parameter 'pageSize' is set
      if (pageSize === undefined || pageSize === null) {
        throw new Error("Missing the required parameter 'pageSize' when calling getServiceNowConfigurationsPage");
      }

      let pathParams = {
      };
      let queryParams = {
        'tenantId': opts['tenantId'],
        'domain': opts['domain'],
        'page': page,
        'pageSize': pageSize
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = ServiceNowConfigurationsPage;
      return this.apiClient.callApi(
        '/incident/api/v8/configurations', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Returns a ServiceNow configurations.
     * @param {Number} page 
     * @param {Number} pageSize 
     * @param {Object} opts Optional parameters
     * @param {String} opts.tenantId 
     * @param {String} opts.domain 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/ServiceNowConfigurationsPage}
     */
    getServiceNowConfigurationsPage(page, pageSize, opts) {
      return this.getServiceNowConfigurationsPageWithHttpInfo(page, pageSize, opts)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * API to patch configure encloud service
     * @param {module:model/IncidentConfigPatch} incidentConfigPatch 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/IncidentConfig} and HTTP response
     */
    patchConfigurationWithHttpInfo(incidentConfigPatch) {
      let postBody = incidentConfigPatch;
      // verify the required parameter 'incidentConfigPatch' is set
      if (incidentConfigPatch === undefined || incidentConfigPatch === null) {
        throw new Error("Missing the required parameter 'incidentConfigPatch' when calling patchConfiguration");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = IncidentConfig;
      return this.apiClient.callApi(
        '/incident/api/v1/config', 'PATCH',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * API to patch configure encloud service
     * @param {module:model/IncidentConfigPatch} incidentConfigPatch 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/IncidentConfig}
     */
    patchConfiguration(incidentConfigPatch) {
      return this.patchConfigurationWithHttpInfo(incidentConfigPatch)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * API to configure encloud service.
     * @param {module:model/IncidentConfigUpdate} incidentConfigUpdate 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/IncidentConfig} and HTTP response
     */
    updateConfigurationWithHttpInfo(incidentConfigUpdate) {
      let postBody = incidentConfigUpdate;
      // verify the required parameter 'incidentConfigUpdate' is set
      if (incidentConfigUpdate === undefined || incidentConfigUpdate === null) {
        throw new Error("Missing the required parameter 'incidentConfigUpdate' when calling updateConfiguration");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = IncidentConfig;
      return this.apiClient.callApi(
        '/incident/api/v1/config', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * API to configure encloud service.
     * @param {module:model/IncidentConfigUpdate} incidentConfigUpdate 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/IncidentConfig}
     */
    updateConfiguration(incidentConfigUpdate) {
      return this.updateConfigurationWithHttpInfo(incidentConfigUpdate)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


}