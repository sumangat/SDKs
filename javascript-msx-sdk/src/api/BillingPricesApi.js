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
import BillingPrice from '../model/BillingPrice';
import BillingPriceCreate from '../model/BillingPriceCreate';
import BillingPriceUpdate from '../model/BillingPriceUpdate';
import BillingPricesPage from '../model/BillingPricesPage';
import Error from '../model/Error';

/**
* BillingPrices service.
* @module api/BillingPricesApi
* @version 1.0.9
*/
export default class BillingPricesApi {

    /**
    * Constructs a new BillingPricesApi. 
    * @alias module:api/BillingPricesApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }



    /**
     * Add price for tenant and event type.
     * Needs MANAGE_PRICES permission to allow for the creation of a price.
     * @param {module:model/BillingPriceCreate} billingPriceCreate 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/BillingPrice} and HTTP response
     */
    addPriceWithHttpInfo(billingPriceCreate) {
      let postBody = billingPriceCreate;
      // verify the required parameter 'billingPriceCreate' is set
      if (billingPriceCreate === undefined || billingPriceCreate === null) {
        throw new Error("Missing the required parameter 'billingPriceCreate' when calling addPrice");
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
      let returnType = BillingPrice;
      return this.apiClient.callApi(
        '/billing/api/v8/prices', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Add price for tenant and event type.
     * Needs MANAGE_PRICES permission to allow for the creation of a price.
     * @param {module:model/BillingPriceCreate} billingPriceCreate 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/BillingPrice}
     */
    addPrice(billingPriceCreate) {
      return this.addPriceWithHttpInfo(billingPriceCreate)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Delete a price.
     * Needs MANAGE_PRICES permission to delete a price.
     * @param {String} id 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing HTTP response
     */
    deletePriceWithHttpInfo(id) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling deletePrice");
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
        '/billing/api/v8/prices/{id}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Delete a price.
     * Needs MANAGE_PRICES permission to delete a price.
     * @param {String} id 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}
     */
    deletePrice(id) {
      return this.deletePriceWithHttpInfo(id)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Get a price.
     * Needs VIEW_PRICES permission to get pricing detail.
     * @param {String} id 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/BillingPrice} and HTTP response
     */
    getPriceWithHttpInfo(id) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getPrice");
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
      let returnType = BillingPrice;
      return this.apiClient.callApi(
        '/billing/api/v8/prices/{id}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Get a price.
     * Needs VIEW_PRICES permission to get pricing detail.
     * @param {String} id 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/BillingPrice}
     */
    getPrice(id) {
      return this.getPriceWithHttpInfo(id)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Retrieve a page of prices.
     * Needs VIEW_PRICES permission to view the pricing details.
     * @param {String} tenantId 
     * @param {Number} page 
     * @param {Number} pageSize 
     * @param {Object} opts Optional parameters
     * @param {String} opts.type 
     * @param {String} opts.subtype 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/BillingPricesPage} and HTTP response
     */
    getPricesPageWithHttpInfo(tenantId, page, pageSize, opts) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'tenantId' is set
      if (tenantId === undefined || tenantId === null) {
        throw new Error("Missing the required parameter 'tenantId' when calling getPricesPage");
      }
      // verify the required parameter 'page' is set
      if (page === undefined || page === null) {
        throw new Error("Missing the required parameter 'page' when calling getPricesPage");
      }
      // verify the required parameter 'pageSize' is set
      if (pageSize === undefined || pageSize === null) {
        throw new Error("Missing the required parameter 'pageSize' when calling getPricesPage");
      }

      let pathParams = {
      };
      let queryParams = {
        'type': opts['type'],
        'subtype': opts['subtype'],
        'tenantId': tenantId,
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
      let returnType = BillingPricesPage;
      return this.apiClient.callApi(
        '/billing/api/v8/prices', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Retrieve a page of prices.
     * Needs VIEW_PRICES permission to view the pricing details.
     * @param {String} tenantId 
     * @param {Number} page 
     * @param {Number} pageSize 
     * @param {Object} opts Optional parameters
     * @param {String} opts.type 
     * @param {String} opts.subtype 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/BillingPricesPage}
     */
    getPricesPage(tenantId, page, pageSize, opts) {
      return this.getPricesPageWithHttpInfo(tenantId, page, pageSize, opts)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Update price for an event type and tenant.
     * Needs MANAGE_PRICES permission to update a pricing detail.
     * @param {String} id 
     * @param {module:model/BillingPriceUpdate} billingPriceUpdate 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/BillingPrice} and HTTP response
     */
    updatePriceWithHttpInfo(id, billingPriceUpdate) {
      let postBody = billingPriceUpdate;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling updatePrice");
      }
      // verify the required parameter 'billingPriceUpdate' is set
      if (billingPriceUpdate === undefined || billingPriceUpdate === null) {
        throw new Error("Missing the required parameter 'billingPriceUpdate' when calling updatePrice");
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
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = BillingPrice;
      return this.apiClient.callApi(
        '/billing/api/v8/prices/{id}', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Update price for an event type and tenant.
     * Needs MANAGE_PRICES permission to update a pricing detail.
     * @param {String} id 
     * @param {module:model/BillingPriceUpdate} billingPriceUpdate 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/BillingPrice}
     */
    updatePrice(id, billingPriceUpdate) {
      return this.updatePriceWithHttpInfo(id, billingPriceUpdate)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


}