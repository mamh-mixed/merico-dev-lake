/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/apache/incubator-devlake/plugins/coding/impl"
	"github.com/apache/incubator-devlake/runner"
	"github.com/spf13/cobra"
)

// PluginEntry Export a variable named PluginEntry for Framework to search and load
var PluginEntry impl.Coding //nolint

// standalone mode for debugging
func main() {
	cmd := &cobra.Command{Use: "coding"}

	// TODO add your cmd flag if necessary
	//depotId := cmd.Flags().Int64P("depotId", "d", 9748559, "depotId")
	//connectionId := cmd.Flags().Int64P("connectionId", "c", 1, "coding connection id")

	cmd.Run = func(cmd *cobra.Command, args []string) {
		runner.DirectRun(cmd, args, PluginEntry, map[string]interface{}{
			"connectionId": 1,
			"depotId":      9748559,
		})
	}
	runner.RunCmd(cmd)
}
