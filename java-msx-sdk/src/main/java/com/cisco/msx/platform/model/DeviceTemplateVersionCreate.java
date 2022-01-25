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
import com.cisco.msx.platform.model.TemplateParameterValidator;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * DeviceTemplateVersionCreate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-01-21T12:39:36.897411-07:00[America/Edmonton]")
public class DeviceTemplateVersionCreate {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_CONFIG_CONTENT = "configContent";
  @SerializedName(SERIALIZED_NAME_CONFIG_CONTENT)
  private String configContent;

  public static final String SERIALIZED_NAME_TEMPLATE_PARAMETER_VALIDATORS = "templateParameterValidators";
  @SerializedName(SERIALIZED_NAME_TEMPLATE_PARAMETER_VALIDATORS)
  private List<TemplateParameterValidator> templateParameterValidators = null;


  public DeviceTemplateVersionCreate name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @ApiModelProperty(required = true, value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public DeviceTemplateVersionCreate configContent(String configContent) {
    
    this.configContent = configContent;
    return this;
  }

   /**
   * Get configContent
   * @return configContent
  **/
  @ApiModelProperty(required = true, value = "")

  public String getConfigContent() {
    return configContent;
  }


  public void setConfigContent(String configContent) {
    this.configContent = configContent;
  }


  public DeviceTemplateVersionCreate templateParameterValidators(List<TemplateParameterValidator> templateParameterValidators) {
    
    this.templateParameterValidators = templateParameterValidators;
    return this;
  }

  public DeviceTemplateVersionCreate addTemplateParameterValidatorsItem(TemplateParameterValidator templateParameterValidatorsItem) {
    if (this.templateParameterValidators == null) {
      this.templateParameterValidators = new ArrayList<>();
    }
    this.templateParameterValidators.add(templateParameterValidatorsItem);
    return this;
  }

   /**
   * Get templateParameterValidators
   * @return templateParameterValidators
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<TemplateParameterValidator> getTemplateParameterValidators() {
    return templateParameterValidators;
  }


  public void setTemplateParameterValidators(List<TemplateParameterValidator> templateParameterValidators) {
    this.templateParameterValidators = templateParameterValidators;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    DeviceTemplateVersionCreate deviceTemplateVersionCreate = (DeviceTemplateVersionCreate) o;
    return Objects.equals(this.name, deviceTemplateVersionCreate.name) &&
        Objects.equals(this.configContent, deviceTemplateVersionCreate.configContent) &&
        Objects.equals(this.templateParameterValidators, deviceTemplateVersionCreate.templateParameterValidators);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, configContent, templateParameterValidators);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DeviceTemplateVersionCreate {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    configContent: ").append(toIndentedString(configContent)).append("\n");
    sb.append("    templateParameterValidators: ").append(toIndentedString(templateParameterValidators)).append("\n");
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
