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

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.JavascriptMsxSdk);
  }
}(this, function(expect, JavascriptMsxSdk) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new JavascriptMsxSdk.BillingCycleAllOf();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('BillingCycleAllOf', function() {
    it('should create an instance of BillingCycleAllOf', function() {
      // uncomment below and update the code to test BillingCycleAllOf
      //var instane = new JavascriptMsxSdk.BillingCycleAllOf();
      //expect(instance).to.be.a(JavascriptMsxSdk.BillingCycleAllOf);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new JavascriptMsxSdk.BillingCycleAllOf();
      //expect(instance).to.be();
    });

    it('should have the property eventId (base name: "eventId")', function() {
      // uncomment below and update the code to test the property eventId
      //var instance = new JavascriptMsxSdk.BillingCycleAllOf();
      //expect(instance).to.be();
    });

    it('should have the property lastBilledOn (base name: "lastBilledOn")', function() {
      // uncomment below and update the code to test the property lastBilledOn
      //var instance = new JavascriptMsxSdk.BillingCycleAllOf();
      //expect(instance).to.be();
    });

    it('should have the property nextBilledOn (base name: "nextBilledOn")', function() {
      // uncomment below and update the code to test the property nextBilledOn
      //var instance = new JavascriptMsxSdk.BillingCycleAllOf();
      //expect(instance).to.be();
    });

    it('should have the property tenantId (base name: "tenantId")', function() {
      // uncomment below and update the code to test the property tenantId
      //var instance = new JavascriptMsxSdk.BillingCycleAllOf();
      //expect(instance).to.be();
    });

  });

}));