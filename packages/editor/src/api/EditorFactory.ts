/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import { EditorInitArgs, KogitoEditorEnvelopeApi } from "./KogitoEditorEnvelopeApi";
import { Editor } from "./Editor";
import { KogitoEditorEnvelopeContextType } from "./KogitoEditorEnvelopeContext";
import { ApiDefinition } from "@kie-tools-core/envelope-bus/dist/api";
import { KogitoEditorChannelApi } from "./KogitoEditorChannelApi";

/**
 * Factory of Editors to be created inside the envelope.
 */
export interface EditorFactory<
  E extends Editor,
  EnvelopeApi extends KogitoEditorEnvelopeApi & ApiDefinition<EnvelopeApi>,
  ChannelApi extends KogitoEditorChannelApi & ApiDefinition<ChannelApi>,
> {
  /**
   * Returns an Editor instance.
   * @param envelopeContext The context to be used by Editor implementation.
   * @param initArgs Initial arguments required for the Editor to initialize itself properly.
   */
  createEditor(
    envelopeContext: KogitoEditorEnvelopeContextType<EnvelopeApi, ChannelApi>,
    initArgs: EditorInitArgs
  ): Promise<E>;
}
