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

import ApiClient from '../ApiClient';

/**
 * The ChangeRequestUpdate model module.
 * @module model/ChangeRequestUpdate
 * @version 1.0.9
 */
class ChangeRequestUpdate {
    /**
     * Constructs a new <code>ChangeRequestUpdate</code>.
     * @alias module:model/ChangeRequestUpdate
     */
    constructor() { 
        
        ChangeRequestUpdate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ChangeRequestUpdate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ChangeRequestUpdate} obj Optional instance to populate.
     * @return {module:model/ChangeRequestUpdate} The populated <code>ChangeRequestUpdate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ChangeRequestUpdate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('attributes')) {
                obj['attributes'] = ApiClient.convertToType(data['attributes'], {'String': Object});
            }
            if (data.hasOwnProperty('tenantId')) {
                obj['tenantId'] = ApiClient.convertToType(data['tenantId'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} name
 */
ChangeRequestUpdate.prototype['name'] = undefined;

/**
 * @member {String} description
 */
ChangeRequestUpdate.prototype['description'] = undefined;

/**
 * @member {Object.<String, Object>} attributes
 */
ChangeRequestUpdate.prototype['attributes'] = undefined;

/**
 * @member {String} tenantId
 */
ChangeRequestUpdate.prototype['tenantId'] = undefined;






export default ChangeRequestUpdate;
