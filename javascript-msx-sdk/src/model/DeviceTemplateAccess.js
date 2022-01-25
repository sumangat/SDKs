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
 * The DeviceTemplateAccess model module.
 * @module model/DeviceTemplateAccess
 * @version 1.0.9
 */
class DeviceTemplateAccess {
    /**
     * Constructs a new <code>DeviceTemplateAccess</code>.
     * @alias module:model/DeviceTemplateAccess
     */
    constructor() { 
        
        DeviceTemplateAccess.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>DeviceTemplateAccess</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/DeviceTemplateAccess} obj Optional instance to populate.
     * @return {module:model/DeviceTemplateAccess} The populated <code>DeviceTemplateAccess</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new DeviceTemplateAccess();

            if (data.hasOwnProperty('tenantIds')) {
                obj['tenantIds'] = ApiClient.convertToType(data['tenantIds'], ['String']);
            }
            if (data.hasOwnProperty('global')) {
                obj['global'] = ApiClient.convertToType(data['global'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * List of tenants to grant access. Mutually exclusive from global flag field.
 * @member {Array.<String>} tenantIds
 */
DeviceTemplateAccess.prototype['tenantIds'] = undefined;

/**
 * Determines if the template is globally accessible. Mutually exclusive from tenant list field.
 * @member {Boolean} global
 */
DeviceTemplateAccess.prototype['global'] = undefined;






export default DeviceTemplateAccess;
