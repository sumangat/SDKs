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
 * The TemplateApplicationCreate model module.
 * @module model/TemplateApplicationCreate
 * @version 1.0.9
 */
class TemplateApplicationCreate {
    /**
     * Constructs a new <code>TemplateApplicationCreate</code>.
     * @alias module:model/TemplateApplicationCreate
     * @param tenantId {String} 
     * @param targetId {String} 
     * @param targetType {String} 
     */
    constructor(tenantId, targetId, targetType) { 
        
        TemplateApplicationCreate.initialize(this, tenantId, targetId, targetType);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, tenantId, targetId, targetType) { 
        obj['tenantId'] = tenantId;
        obj['targetId'] = targetId;
        obj['targetType'] = targetType;
    }

    /**
     * Constructs a <code>TemplateApplicationCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TemplateApplicationCreate} obj Optional instance to populate.
     * @return {module:model/TemplateApplicationCreate} The populated <code>TemplateApplicationCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TemplateApplicationCreate();

            if (data.hasOwnProperty('tenantId')) {
                obj['tenantId'] = ApiClient.convertToType(data['tenantId'], 'String');
            }
            if (data.hasOwnProperty('targetId')) {
                obj['targetId'] = ApiClient.convertToType(data['targetId'], 'String');
            }
            if (data.hasOwnProperty('targetType')) {
                obj['targetType'] = ApiClient.convertToType(data['targetType'], 'String');
            }
            if (data.hasOwnProperty('parameters')) {
                obj['parameters'] = ApiClient.convertToType(data['parameters'], {'String': 'String'});
            }
        }
        return obj;
    }


}

/**
 * @member {String} tenantId
 */
TemplateApplicationCreate.prototype['tenantId'] = undefined;

/**
 * @member {String} targetId
 */
TemplateApplicationCreate.prototype['targetId'] = undefined;

/**
 * @member {String} targetType
 */
TemplateApplicationCreate.prototype['targetType'] = undefined;

/**
 * @member {Object.<String, String>} parameters
 */
TemplateApplicationCreate.prototype['parameters'] = undefined;






export default TemplateApplicationCreate;
