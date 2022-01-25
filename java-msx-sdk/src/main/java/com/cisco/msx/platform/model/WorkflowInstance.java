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
import com.cisco.msx.platform.model.WorkflowAction;
import com.cisco.msx.platform.model.WorkflowFooter;
import com.cisco.msx.platform.model.WorkflowInstanceAllOf;
import com.cisco.msx.platform.model.WorkflowVariable;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * WorkflowInstance
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-01-21T12:39:36.897411-07:00[America/Edmonton]")
public class WorkflowInstance {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private String id;

  public static final String SERIALIZED_NAME_DEFINITION_ID = "definition_id";
  @SerializedName(SERIALIZED_NAME_DEFINITION_ID)
  private String definitionId;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_SCHEMA_ID = "schema_id";
  @SerializedName(SERIALIZED_NAME_SCHEMA_ID)
  private String schemaId;

  public static final String SERIALIZED_NAME_VERSION = "version";
  @SerializedName(SERIALIZED_NAME_VERSION)
  private String version;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;

  public static final String SERIALIZED_NAME_BASE_TYPE = "base_type";
  @SerializedName(SERIALIZED_NAME_BASE_TYPE)
  private String baseType;

  public static final String SERIALIZED_NAME_PROPERTIES = "properties";
  @SerializedName(SERIALIZED_NAME_PROPERTIES)
  private Map<String, Object> properties = null;

  public static final String SERIALIZED_NAME_ACTIONS = "actions";
  @SerializedName(SERIALIZED_NAME_ACTIONS)
  private List<WorkflowAction> actions = null;

  public static final String SERIALIZED_NAME_VARIABLES = "variables";
  @SerializedName(SERIALIZED_NAME_VARIABLES)
  private List<WorkflowVariable> variables = null;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private Map<String, Object> status = null;

  public static final String SERIALIZED_NAME_STARTED_ON = "started_on";
  @SerializedName(SERIALIZED_NAME_STARTED_ON)
  private String startedOn;

  public static final String SERIALIZED_NAME_ENDED_ON = "ended_on";
  @SerializedName(SERIALIZED_NAME_ENDED_ON)
  private String endedOn;

  public static final String SERIALIZED_NAME_CREATED_ON = "created_on";
  @SerializedName(SERIALIZED_NAME_CREATED_ON)
  private String createdOn;

  public static final String SERIALIZED_NAME_CREATED_BY = "created_by";
  @SerializedName(SERIALIZED_NAME_CREATED_BY)
  private String createdBy;

  public static final String SERIALIZED_NAME_UPDATED_ON = "updated_on";
  @SerializedName(SERIALIZED_NAME_UPDATED_ON)
  private String updatedOn;

  public static final String SERIALIZED_NAME_UPDATED_BY = "updated_by";
  @SerializedName(SERIALIZED_NAME_UPDATED_BY)
  private String updatedBy;

  public static final String SERIALIZED_NAME_OWNER = "owner";
  @SerializedName(SERIALIZED_NAME_OWNER)
  private String owner;

  public static final String SERIALIZED_NAME_UNIQUE_NAME = "unique_name";
  @SerializedName(SERIALIZED_NAME_UNIQUE_NAME)
  private String uniqueName;


  public WorkflowInstance id(String id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getId() {
    return id;
  }


  public void setId(String id) {
    this.id = id;
  }


  public WorkflowInstance definitionId(String definitionId) {
    
    this.definitionId = definitionId;
    return this;
  }

   /**
   * Get definitionId
   * @return definitionId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDefinitionId() {
    return definitionId;
  }


  public void setDefinitionId(String definitionId) {
    this.definitionId = definitionId;
  }


  public WorkflowInstance name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public WorkflowInstance schemaId(String schemaId) {
    
    this.schemaId = schemaId;
    return this;
  }

   /**
   * Get schemaId
   * @return schemaId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSchemaId() {
    return schemaId;
  }


  public void setSchemaId(String schemaId) {
    this.schemaId = schemaId;
  }


  public WorkflowInstance version(String version) {
    
    this.version = version;
    return this;
  }

   /**
   * Get version
   * @return version
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVersion() {
    return version;
  }


  public void setVersion(String version) {
    this.version = version;
  }


  public WorkflowInstance type(String type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getType() {
    return type;
  }


  public void setType(String type) {
    this.type = type;
  }


  public WorkflowInstance baseType(String baseType) {
    
    this.baseType = baseType;
    return this;
  }

   /**
   * Get baseType
   * @return baseType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getBaseType() {
    return baseType;
  }


  public void setBaseType(String baseType) {
    this.baseType = baseType;
  }


  public WorkflowInstance properties(Map<String, Object> properties) {
    
    this.properties = properties;
    return this;
  }

  public WorkflowInstance putPropertiesItem(String key, Object propertiesItem) {
    if (this.properties == null) {
      this.properties = new HashMap<>();
    }
    this.properties.put(key, propertiesItem);
    return this;
  }

   /**
   * Get properties
   * @return properties
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, Object> getProperties() {
    return properties;
  }


  public void setProperties(Map<String, Object> properties) {
    this.properties = properties;
  }


  public WorkflowInstance actions(List<WorkflowAction> actions) {
    
    this.actions = actions;
    return this;
  }

  public WorkflowInstance addActionsItem(WorkflowAction actionsItem) {
    if (this.actions == null) {
      this.actions = new ArrayList<>();
    }
    this.actions.add(actionsItem);
    return this;
  }

   /**
   * Get actions
   * @return actions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<WorkflowAction> getActions() {
    return actions;
  }


  public void setActions(List<WorkflowAction> actions) {
    this.actions = actions;
  }


  public WorkflowInstance variables(List<WorkflowVariable> variables) {
    
    this.variables = variables;
    return this;
  }

  public WorkflowInstance addVariablesItem(WorkflowVariable variablesItem) {
    if (this.variables == null) {
      this.variables = new ArrayList<>();
    }
    this.variables.add(variablesItem);
    return this;
  }

   /**
   * Get variables
   * @return variables
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<WorkflowVariable> getVariables() {
    return variables;
  }


  public void setVariables(List<WorkflowVariable> variables) {
    this.variables = variables;
  }


  public WorkflowInstance status(Map<String, Object> status) {
    
    this.status = status;
    return this;
  }

  public WorkflowInstance putStatusItem(String key, Object statusItem) {
    if (this.status == null) {
      this.status = new HashMap<>();
    }
    this.status.put(key, statusItem);
    return this;
  }

   /**
   * Get status
   * @return status
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, Object> getStatus() {
    return status;
  }


  public void setStatus(Map<String, Object> status) {
    this.status = status;
  }


  public WorkflowInstance startedOn(String startedOn) {
    
    this.startedOn = startedOn;
    return this;
  }

   /**
   * Get startedOn
   * @return startedOn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStartedOn() {
    return startedOn;
  }


  public void setStartedOn(String startedOn) {
    this.startedOn = startedOn;
  }


  public WorkflowInstance endedOn(String endedOn) {
    
    this.endedOn = endedOn;
    return this;
  }

   /**
   * Get endedOn
   * @return endedOn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getEndedOn() {
    return endedOn;
  }


  public void setEndedOn(String endedOn) {
    this.endedOn = endedOn;
  }


  public WorkflowInstance createdOn(String createdOn) {
    
    this.createdOn = createdOn;
    return this;
  }

   /**
   * Get createdOn
   * @return createdOn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCreatedOn() {
    return createdOn;
  }


  public void setCreatedOn(String createdOn) {
    this.createdOn = createdOn;
  }


  public WorkflowInstance createdBy(String createdBy) {
    
    this.createdBy = createdBy;
    return this;
  }

   /**
   * Get createdBy
   * @return createdBy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCreatedBy() {
    return createdBy;
  }


  public void setCreatedBy(String createdBy) {
    this.createdBy = createdBy;
  }


  public WorkflowInstance updatedOn(String updatedOn) {
    
    this.updatedOn = updatedOn;
    return this;
  }

   /**
   * Get updatedOn
   * @return updatedOn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUpdatedOn() {
    return updatedOn;
  }


  public void setUpdatedOn(String updatedOn) {
    this.updatedOn = updatedOn;
  }


  public WorkflowInstance updatedBy(String updatedBy) {
    
    this.updatedBy = updatedBy;
    return this;
  }

   /**
   * Get updatedBy
   * @return updatedBy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUpdatedBy() {
    return updatedBy;
  }


  public void setUpdatedBy(String updatedBy) {
    this.updatedBy = updatedBy;
  }


  public WorkflowInstance owner(String owner) {
    
    this.owner = owner;
    return this;
  }

   /**
   * Get owner
   * @return owner
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getOwner() {
    return owner;
  }


  public void setOwner(String owner) {
    this.owner = owner;
  }


  public WorkflowInstance uniqueName(String uniqueName) {
    
    this.uniqueName = uniqueName;
    return this;
  }

   /**
   * Get uniqueName
   * @return uniqueName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUniqueName() {
    return uniqueName;
  }


  public void setUniqueName(String uniqueName) {
    this.uniqueName = uniqueName;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    WorkflowInstance workflowInstance = (WorkflowInstance) o;
    return Objects.equals(this.id, workflowInstance.id) &&
        Objects.equals(this.definitionId, workflowInstance.definitionId) &&
        Objects.equals(this.name, workflowInstance.name) &&
        Objects.equals(this.schemaId, workflowInstance.schemaId) &&
        Objects.equals(this.version, workflowInstance.version) &&
        Objects.equals(this.type, workflowInstance.type) &&
        Objects.equals(this.baseType, workflowInstance.baseType) &&
        Objects.equals(this.properties, workflowInstance.properties) &&
        Objects.equals(this.actions, workflowInstance.actions) &&
        Objects.equals(this.variables, workflowInstance.variables) &&
        Objects.equals(this.status, workflowInstance.status) &&
        Objects.equals(this.startedOn, workflowInstance.startedOn) &&
        Objects.equals(this.endedOn, workflowInstance.endedOn) &&
        Objects.equals(this.createdOn, workflowInstance.createdOn) &&
        Objects.equals(this.createdBy, workflowInstance.createdBy) &&
        Objects.equals(this.updatedOn, workflowInstance.updatedOn) &&
        Objects.equals(this.updatedBy, workflowInstance.updatedBy) &&
        Objects.equals(this.owner, workflowInstance.owner) &&
        Objects.equals(this.uniqueName, workflowInstance.uniqueName);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, definitionId, name, schemaId, version, type, baseType, properties, actions, variables, status, startedOn, endedOn, createdOn, createdBy, updatedOn, updatedBy, owner, uniqueName);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class WorkflowInstance {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    definitionId: ").append(toIndentedString(definitionId)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    schemaId: ").append(toIndentedString(schemaId)).append("\n");
    sb.append("    version: ").append(toIndentedString(version)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    baseType: ").append(toIndentedString(baseType)).append("\n");
    sb.append("    properties: ").append(toIndentedString(properties)).append("\n");
    sb.append("    actions: ").append(toIndentedString(actions)).append("\n");
    sb.append("    variables: ").append(toIndentedString(variables)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    startedOn: ").append(toIndentedString(startedOn)).append("\n");
    sb.append("    endedOn: ").append(toIndentedString(endedOn)).append("\n");
    sb.append("    createdOn: ").append(toIndentedString(createdOn)).append("\n");
    sb.append("    createdBy: ").append(toIndentedString(createdBy)).append("\n");
    sb.append("    updatedOn: ").append(toIndentedString(updatedOn)).append("\n");
    sb.append("    updatedBy: ").append(toIndentedString(updatedBy)).append("\n");
    sb.append("    owner: ").append(toIndentedString(owner)).append("\n");
    sb.append("    uniqueName: ").append(toIndentedString(uniqueName)).append("\n");
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
