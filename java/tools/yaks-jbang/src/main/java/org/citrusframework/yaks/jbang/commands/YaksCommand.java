/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.citrusframework.yaks.jbang.commands;

import java.io.File;
import java.util.Stack;
import java.util.concurrent.Callable;

import org.citrusframework.yaks.jbang.YaksJBangMain;
import picocli.CommandLine;
import picocli.CommandLine.IParameterConsumer;
import picocli.CommandLine.Model.ArgSpec;
import picocli.CommandLine.Model.CommandSpec;
import picocli.CommandLine.ParameterException;

public abstract class YaksCommand implements Callable<Integer> {

    @CommandLine.Spec
    CommandSpec spec;

    private final YaksJBangMain main;
    private File yaksDir;

    //CHECKSTYLE:OFF
    @CommandLine.Option(names = { "-h", "--help" }, usageHelp = true, description = "Display the help and sub-commands")
    private boolean helpRequested = false;
    //CHECKSTYLE:ON

    public YaksCommand(YaksJBangMain main) {
        this.main = main;
    }

    public YaksJBangMain getMain() {
        return main;
    }

    public File getStatusFile(String pid) {
        if (yaksDir == null) {
            yaksDir = new File(System.getProperty("user.home"), ".yaks");
        }
        return new File(yaksDir, pid + "-status.json");
    }

    public File getOutputFile(String pid) {
        if (yaksDir == null) {
            yaksDir = new File(System.getProperty("user.home"), ".yaks");
        }
        return new File(yaksDir, pid + "-output.json");
    }

    protected abstract static class ParameterConsumer<T> implements IParameterConsumer {

        @Override
        public void consumeParameters(Stack<String> args, ArgSpec argSpec, CommandSpec cmdSpec) {
            if (args.isEmpty()) {
                throw new ParameterException(cmdSpec.commandLine(), "Error: missing required parameter");
            }
            T cmd = (T) cmdSpec.userObject();
            doConsumeParameters(args, cmd);
        }

        protected abstract void doConsumeParameters(Stack<String> args, T cmd);
    }

}
