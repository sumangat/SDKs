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
import Product from './Product';

/**
 * The ProductsPageAllOf model module.
 * @module model/ProductsPageAllOf
 * @version 1.0.9
 */
class ProductsPageAllOf {
    /**
     * Constructs a new <code>ProductsPageAllOf</code>.
     * @alias module:model/ProductsPageAllOf
     */
    constructor() { 
        
        ProductsPageAllOf.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ProductsPageAllOf</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ProductsPageAllOf} obj Optional instance to populate.
     * @return {module:model/ProductsPageAllOf} The populated <code>ProductsPageAllOf</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ProductsPageAllOf();

            if (data.hasOwnProperty('contents')) {
                obj['contents'] = ApiClient.convertToType(data['contents'], [Product]);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/Product>} contents
 */
ProductsPageAllOf.prototype['contents'] = undefined;






export default ProductsPageAllOf;
