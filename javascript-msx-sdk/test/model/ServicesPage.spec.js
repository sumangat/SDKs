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
    instance = new JavascriptMsxSdk.ServicesPage();
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

  describe('ServicesPage', function() {
    it('should create an instance of ServicesPage', function() {
      // uncomment below and update the code to test ServicesPage
      //var instane = new JavascriptMsxSdk.ServicesPage();
      //expect(instance).to.be.a(JavascriptMsxSdk.ServicesPage);
    });

    it('should have the property page (base name: "page")', function() {
      // uncomment below and update the code to test the property page
      //var instance = new JavascriptMsxSdk.ServicesPage();
      //expect(instance).to.be();
    });

    it('should have the property pageSize (base name: "pageSize")', function() {
      // uncomment below and update the code to test the property pageSize
      //var instance = new JavascriptMsxSdk.ServicesPage();
      //expect(instance).to.be();
    });

    it('should have the property totalItems (base name: "totalItems")', function() {
      // uncomment below and update the code to test the property totalItems
      //var instance = new JavascriptMsxSdk.ServicesPage();
      //expect(instance).to.be();
    });

    it('should have the property hasNext (base name: "hasNext")', function() {
      // uncomment below and update the code to test the property hasNext
      //var instance = new JavascriptMsxSdk.ServicesPage();
      //expect(instance).to.be();
    });

    it('should have the property hasPrevious (base name: "hasPrevious")', function() {
      // uncomment below and update the code to test the property hasPrevious
      //var instance = new JavascriptMsxSdk.ServicesPage();
      //expect(instance).to.be();
    });

    it('should have the property sortBy (base name: "sortBy")', function() {
      // uncomment below and update the code to test the property sortBy
      //var instance = new JavascriptMsxSdk.ServicesPage();
      //expect(instance).to.be();
    });

    it('should have the property sortOrder (base name: "sortOrder")', function() {
      // uncomment below and update the code to test the property sortOrder
      //var instance = new JavascriptMsxSdk.ServicesPage();
      //expect(instance).to.be();
    });

    it('should have the property contents (base name: "contents")', function() {
      // uncomment below and update the code to test the property contents
      //var instance = new JavascriptMsxSdk.ServicesPage();
      //expect(instance).to.be();
    });

  });

}));