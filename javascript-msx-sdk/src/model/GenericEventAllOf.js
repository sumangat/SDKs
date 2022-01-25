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
 * The GenericEventAllOf model module.
 * @module model/GenericEventAllOf
 * @version 1.0.9
 */
class GenericEventAllOf {
    /**
     * Constructs a new <code>GenericEventAllOf</code>.
     * @alias module:model/GenericEventAllOf
     */
    constructor() { 
        
        GenericEventAllOf.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>GenericEventAllOf</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/GenericEventAllOf} obj Optional instance to populate.
     * @return {module:model/GenericEventAllOf} The populated <code>GenericEventAllOf</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new GenericEventAllOf();

            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('time')) {
                obj['time'] = ApiClient.convertToType(data['time'], 'Date');
            }
        }
        return obj;
    }


}

/**
 * @member {String} type
 */
GenericEventAllOf.prototype['type'] = undefined;

/**
 * @member {Date} time
 */
GenericEventAllOf.prototype['time'] = undefined;






export default GenericEventAllOf;
