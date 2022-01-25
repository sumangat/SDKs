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
import com.cisco.msx.platform.model.WorkflowAccessMetaType;
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
 * WorkflowAccessMeta
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-01-21T12:39:36.897411-07:00[America/Edmonton]")
public class WorkflowAccessMeta {
  public static final String SERIALIZED_NAME_ADAPTER = "adapter";
  @SerializedName(SERIALIZED_NAME_ADAPTER)
  private WorkflowAccessMetaType adapter;

  public static final String SERIALIZED_NAME_RUNTIME_USERS = "runtime_users";
  @SerializedName(SERIALIZED_NAME_RUNTIME_USERS)
  private List<WorkflowAccessMetaType> runtimeUsers = null;

  public static final String SERIALIZED_NAME_TARGETS = "targets";
  @SerializedName(SERIALIZED_NAME_TARGETS)
  private List<WorkflowAccessMetaType> targets = null;

  public static final String SERIALIZED_NAME_IS_INTEGRATION = "is_integration";
  @SerializedName(SERIALIZED_NAME_IS_INTEGRATION)
  private Boolean isIntegration;

  public static final String SERIALIZED_NAME_IS_INTERNAL = "is_internal";
  @SerializedName(SERIALIZED_NAME_IS_INTERNAL)
  private Boolean isInternal;


  public WorkflowAccessMeta adapter(WorkflowAccessMetaType adapter) {
    
    this.adapter = adapter;
    return this;
  }

   /**
   * Get adapter
   * @return adapter
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public WorkflowAccessMetaType getAdapter() {
    return adapter;
  }


  public void setAdapter(WorkflowAccessMetaType adapter) {
    this.adapter = adapter;
  }


  public WorkflowAccessMeta runtimeUsers(List<WorkflowAccessMetaType> runtimeUsers) {
    
    this.runtimeUsers = runtimeUsers;
    return this;
  }

  public WorkflowAccessMeta addRuntimeUsersItem(WorkflowAccessMetaType runtimeUsersItem) {
    if (this.runtimeUsers == null) {
      this.runtimeUsers = new ArrayList<>();
    }
    this.runtimeUsers.add(runtimeUsersItem);
    return this;
  }

   /**
   * Get runtimeUsers
   * @return runtimeUsers
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<WorkflowAccessMetaType> getRuntimeUsers() {
    return runtimeUsers;
  }


  public void setRuntimeUsers(List<WorkflowAccessMetaType> runtimeUsers) {
    this.runtimeUsers = runtimeUsers;
  }


  public WorkflowAccessMeta targets(List<WorkflowAccessMetaType> targets) {
    
    this.targets = targets;
    return this;
  }

  public WorkflowAccessMeta addTargetsItem(WorkflowAccessMetaType targetsItem) {
    if (this.targets == null) {
      this.targets = new ArrayList<>();
    }
    this.targets.add(targetsItem);
    return this;
  }

   /**
   * Get targets
   * @return targets
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<WorkflowAccessMetaType> getTargets() {
    return targets;
  }


  public void setTargets(List<WorkflowAccessMetaType> targets) {
    this.targets = targets;
  }


  public WorkflowAccessMeta isIntegration(Boolean isIntegration) {
    
    this.isIntegration = isIntegration;
    return this;
  }

   /**
   * Get isIntegration
   * @return isIntegration
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getIsIntegration() {
    return isIntegration;
  }


  public void setIsIntegration(Boolean isIntegration) {
    this.isIntegration = isIntegration;
  }


  public WorkflowAccessMeta isInternal(Boolean isInternal) {
    
    this.isInternal = isInternal;
    return this;
  }

   /**
   * Get isInternal
   * @return isInternal
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getIsInternal() {
    return isInternal;
  }


  public void setIsInternal(Boolean isInternal) {
    this.isInternal = isInternal;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    WorkflowAccessMeta workflowAccessMeta = (WorkflowAccessMeta) o;
    return Objects.equals(this.adapter, workflowAccessMeta.adapter) &&
        Objects.equals(this.runtimeUsers, workflowAccessMeta.runtimeUsers) &&
        Objects.equals(this.targets, workflowAccessMeta.targets) &&
        Objects.equals(this.isIntegration, workflowAccessMeta.isIntegration) &&
        Objects.equals(this.isInternal, workflowAccessMeta.isInternal);
  }

  @Override
  public int hashCode() {
    return Objects.hash(adapter, runtimeUsers, targets, isIntegration, isInternal);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class WorkflowAccessMeta {\n");
    sb.append("    adapter: ").append(toIndentedString(adapter)).append("\n");
    sb.append("    runtimeUsers: ").append(toIndentedString(runtimeUsers)).append("\n");
    sb.append("    targets: ").append(toIndentedString(targets)).append("\n");
    sb.append("    isIntegration: ").append(toIndentedString(isIntegration)).append("\n");
    sb.append("    isInternal: ").append(toIndentedString(isInternal)).append("\n");
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
