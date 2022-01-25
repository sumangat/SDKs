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
import TemplateApplicationAllOf from './TemplateApplicationAllOf';
import TemplateApplicationCreate from './TemplateApplicationCreate';
import TemplateStatus from './TemplateStatus';
import TemplateStatusMeta from './TemplateStatusMeta';

/**
 * The TemplateApplication model module.
 * @module model/TemplateApplication
 * @version 1.0.9
 */
class TemplateApplication {
    /**
     * Constructs a new <code>TemplateApplication</code>.
     * @alias module:model/TemplateApplication
     * @implements module:model/TemplateApplicationAllOf
     * @implements module:model/TemplateApplicationCreate
     * @implements module:model/TemplateStatusMeta
     * @param tenantId {String} 
     * @param targetId {String} 
     * @param targetType {String} 
     */
    constructor(tenantId, targetId, targetType) { 
        TemplateApplicationAllOf.initialize(this);TemplateApplicationCreate.initialize(this, tenantId, targetId, targetType);TemplateStatusMeta.initialize(this);
        TemplateApplication.initialize(this, tenantId, targetId, targetType);
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
     * Constructs a <code>TemplateApplication</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TemplateApplication} obj Optional instance to populate.
     * @return {module:model/TemplateApplication} The populated <code>TemplateApplication</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TemplateApplication();
            TemplateApplicationAllOf.constructFromObject(data, obj);
            TemplateApplicationCreate.constructFromObject(data, obj);
            TemplateStatusMeta.constructFromObject(data, obj);

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'String');
            }
            if (data.hasOwnProperty('templateId')) {
                obj['templateId'] = ApiClient.convertToType(data['templateId'], 'String');
            }
            if (data.hasOwnProperty('templateName')) {
                obj['templateName'] = ApiClient.convertToType(data['templateName'], 'String');
            }
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
            if (data.hasOwnProperty('status')) {
                obj['status'] = TemplateStatus.constructFromObject(data['status']);
            }
            if (data.hasOwnProperty('statusDetails')) {
                obj['statusDetails'] = ApiClient.convertToType(data['statusDetails'], 'String');
            }
            if (data.hasOwnProperty('createdOn')) {
                obj['createdOn'] = ApiClient.convertToType(data['createdOn'], 'Date');
            }
            if (data.hasOwnProperty('createdBy')) {
                obj['createdBy'] = ApiClient.convertToType(data['createdBy'], 'String');
            }
            if (data.hasOwnProperty('modifiedOn')) {
                obj['modifiedOn'] = ApiClient.convertToType(data['modifiedOn'], 'Date');
            }
            if (data.hasOwnProperty('modifiedBy')) {
                obj['modifiedBy'] = ApiClient.convertToType(data['modifiedBy'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} id
 */
TemplateApplication.prototype['id'] = undefined;

/**
 * @member {String} templateId
 */
TemplateApplication.prototype['templateId'] = undefined;

/**
 * @member {String} templateName
 */
TemplateApplication.prototype['templateName'] = undefined;

/**
 * @member {String} tenantId
 */
TemplateApplication.prototype['tenantId'] = undefined;

/**
 * @member {String} targetId
 */
TemplateApplication.prototype['targetId'] = undefined;

/**
 * @member {String} targetType
 */
TemplateApplication.prototype['targetType'] = undefined;

/**
 * @member {Object.<String, String>} parameters
 */
TemplateApplication.prototype['parameters'] = undefined;

/**
 * @member {module:model/TemplateStatus} status
 */
TemplateApplication.prototype['status'] = undefined;

/**
 * @member {String} statusDetails
 */
TemplateApplication.prototype['statusDetails'] = undefined;

/**
 * @member {Date} createdOn
 */
TemplateApplication.prototype['createdOn'] = undefined;

/**
 * @member {String} createdBy
 */
TemplateApplication.prototype['createdBy'] = undefined;

/**
 * @member {Date} modifiedOn
 */
TemplateApplication.prototype['modifiedOn'] = undefined;

/**
 * @member {String} modifiedBy
 */
TemplateApplication.prototype['modifiedBy'] = undefined;


// Implement TemplateApplicationAllOf interface:
/**
 * @member {String} id
 */
TemplateApplicationAllOf.prototype['id'] = undefined;
/**
 * @member {String} templateId
 */
TemplateApplicationAllOf.prototype['templateId'] = undefined;
/**
 * @member {String} templateName
 */
TemplateApplicationAllOf.prototype['templateName'] = undefined;
// Implement TemplateApplicationCreate interface:
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
// Implement TemplateStatusMeta interface:
/**
 * @member {module:model/TemplateStatus} status
 */
TemplateStatusMeta.prototype['status'] = undefined;
/**
 * @member {String} statusDetails
 */
TemplateStatusMeta.prototype['statusDetails'] = undefined;
/**
 * @member {Date} createdOn
 */
TemplateStatusMeta.prototype['createdOn'] = undefined;
/**
 * @member {String} createdBy
 */
TemplateStatusMeta.prototype['createdBy'] = undefined;
/**
 * @member {Date} modifiedOn
 */
TemplateStatusMeta.prototype['modifiedOn'] = undefined;
/**
 * @member {String} modifiedBy
 */
TemplateStatusMeta.prototype['modifiedBy'] = undefined;




export default TemplateApplication;
