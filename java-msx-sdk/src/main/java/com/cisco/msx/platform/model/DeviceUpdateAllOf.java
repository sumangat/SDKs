/*
 * MSX SDK
 * MSX SDK client.
 *
 * The version of the OpenAPI document: 1.0.9
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package com.cisco.msx.platform.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * DeviceUpdateAllOf
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-01-21T12:39:36.897411-07:00[America/Edmonton]")
public class DeviceUpdateAllOf {
  public static final String SERIALIZED_NAME_SERVICE_TYPE = "serviceType";
  @SerializedName(SERIALIZED_NAME_SERVICE_TYPE)
  private String serviceType;

  public static final String SERIALIZED_NAME_TAGS = "tags";
  @SerializedName(SERIALIZED_NAME_TAGS)
  private Map<String, String> tags = null;

  public static final String SERIALIZED_NAME_MANAGED = "managed";
  @SerializedName(SERIALIZED_NAME_MANAGED)
  private Boolean managed = false;

  public static final String SERIALIZED_NAME_ONBOARD_TYPE = "onboardType";
  @SerializedName(SERIALIZED_NAME_ONBOARD_TYPE)
  private String onboardType;

  public static final String SERIALIZED_NAME_ONBOARD_INFORMATION = "onboardInformation";
  @SerializedName(SERIALIZED_NAME_ONBOARD_INFORMATION)
  private Map<String, Object> onboardInformation = null;

  public static final String SERIALIZED_NAME_ATTRIBUTES = "attributes";
  @SerializedName(SERIALIZED_NAME_ATTRIBUTES)
  private Map<String, Object> attributes = null;


  public DeviceUpdateAllOf serviceType(String serviceType) {
    
    this.serviceType = serviceType;
    return this;
  }

   /**
   * Get serviceType
   * @return serviceType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getServiceType() {
    return serviceType;
  }


  public void setServiceType(String serviceType) {
    this.serviceType = serviceType;
  }


  public DeviceUpdateAllOf tags(Map<String, String> tags) {
    
    this.tags = tags;
    return this;
  }

  public DeviceUpdateAllOf putTagsItem(String key, String tagsItem) {
    if (this.tags == null) {
      this.tags = new HashMap<>();
    }
    this.tags.put(key, tagsItem);
    return this;
  }

   /**
   * Get tags
   * @return tags
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, String> getTags() {
    return tags;
  }


  public void setTags(Map<String, String> tags) {
    this.tags = tags;
  }


  public DeviceUpdateAllOf managed(Boolean managed) {
    
    this.managed = managed;
    return this;
  }

   /**
   * Get managed
   * @return managed
  **/
  @ApiModelProperty(required = true, value = "")

  public Boolean getManaged() {
    return managed;
  }


  public void setManaged(Boolean managed) {
    this.managed = managed;
  }


  public DeviceUpdateAllOf onboardType(String onboardType) {
    
    this.onboardType = onboardType;
    return this;
  }

   /**
   * Get onboardType
   * @return onboardType
  **/
  @ApiModelProperty(required = true, value = "")

  public String getOnboardType() {
    return onboardType;
  }


  public void setOnboardType(String onboardType) {
    this.onboardType = onboardType;
  }


  public DeviceUpdateAllOf onboardInformation(Map<String, Object> onboardInformation) {
    
    this.onboardInformation = onboardInformation;
    return this;
  }

  public DeviceUpdateAllOf putOnboardInformationItem(String key, Object onboardInformationItem) {
    if (this.onboardInformation == null) {
      this.onboardInformation = new HashMap<>();
    }
    this.onboardInformation.put(key, onboardInformationItem);
    return this;
  }

   /**
   * Get onboardInformation
   * @return onboardInformation
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, Object> getOnboardInformation() {
    return onboardInformation;
  }


  public void setOnboardInformation(Map<String, Object> onboardInformation) {
    this.onboardInformation = onboardInformation;
  }


  public DeviceUpdateAllOf attributes(Map<String, Object> attributes) {
    
    this.attributes = attributes;
    return this;
  }

  public DeviceUpdateAllOf putAttributesItem(String key, Object attributesItem) {
    if (this.attributes == null) {
      this.attributes = new HashMap<>();
    }
    this.attributes.put(key, attributesItem);
    return this;
  }

   /**
   * Get attributes
   * @return attributes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, Object> getAttributes() {
    return attributes;
  }


  public void setAttributes(Map<String, Object> attributes) {
    this.attributes = attributes;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    DeviceUpdateAllOf deviceUpdateAllOf = (DeviceUpdateAllOf) o;
    return Objects.equals(this.serviceType, deviceUpdateAllOf.serviceType) &&
        Objects.equals(this.tags, deviceUpdateAllOf.tags) &&
        Objects.equals(this.managed, deviceUpdateAllOf.managed) &&
        Objects.equals(this.onboardType, deviceUpdateAllOf.onboardType) &&
        Objects.equals(this.onboardInformation, deviceUpdateAllOf.onboardInformation) &&
        Objects.equals(this.attributes, deviceUpdateAllOf.attributes);
  }

  @Override
  public int hashCode() {
    return Objects.hash(serviceType, tags, managed, onboardType, onboardInformation, attributes);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DeviceUpdateAllOf {\n");
    sb.append("    serviceType: ").append(toIndentedString(serviceType)).append("\n");
    sb.append("    tags: ").append(toIndentedString(tags)).append("\n");
    sb.append("    managed: ").append(toIndentedString(managed)).append("\n");
    sb.append("    onboardType: ").append(toIndentedString(onboardType)).append("\n");
    sb.append("    onboardInformation: ").append(toIndentedString(onboardInformation)).append("\n");
    sb.append("    attributes: ").append(toIndentedString(attributes)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

