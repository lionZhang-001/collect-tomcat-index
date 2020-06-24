/*
@Time : 2020/6/19 20:57
@Author : zhangyongyue
@File : basicStructure
@Software: GoLand
*/
package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
)

var port1, port2 string
var catalina1, name1, name2 string
var catalina2 string
var attr1 []string = make([]string, 5)
var attr2 []string = make([]string, 30)
var logs *logrus.Logger = logrus.New()
var logsFile *os.File

var index2Mysql map[string]string = make(map[string]string)

func init() {
	//日志
	//logs = logrus.New()
	logs.SetFormatter(&logrus.TextFormatter{})
	logsFile, err := createLogsFile()
	if err != nil {
		fmt.Printf("creating logs file ", logsFile.Name(), " went wrong :", err)
		return
	}
	logs.SetOutput(logsFile)

	//读取connector port
	data, err := ioutil.ReadFile("../config/connector_port")
	if err != nil {
		logs.Info(err, " , unable to get connector information")
	}
	ports := strings.Split(string(data), "|")
	port1 := ports[0]
	port2 := ports[1]

	catalina1 = "Catalina:type=GlobalRequestProcessor,name="
	catalina2 = "Catalina:type=Connector,port="
	name1 = "\"ajp-bio-" + port1 + "\""
	name2 = "\"http-bio-" + port2 + "\""
	attr1 = []string{"bytesReceived", "bytesSent", "errorCount", "processingTime", "requestCount", "maxTime", "modelerType"}
	attr2 = []string{"acceptCount", "allowTrace", "connectionLinger", "connectionTimeout", "enableLookups", "executorName", "keepAliveTimeout", "localPort", "maxHeaderCount", "maxKeepAliveRequests", "maxParameterCount",
		"maxPostSize", "maxSavePostSize", "maxSwallowSize", "maxThreads", "minSpareThreads", "port", "processorCache", "protocol", "protocolHandlerClassName", "proxyPort", "redirectPort", "scheme", "secure", "stateName", "tcpNoDelay", "threadPriority", "useBodyEncodingForURI", "useIPVHosts", "xpoweredBy"}

	index2Mysql = map[string]string{
		catalina1 + name1 + ",attr=" + attr1[0]:                              "ajp_bytes_received",
		catalina1 + name1 + ",attr=" + attr1[1]:                              "ajp_bytes_sent",
		catalina1 + name1 + ",attr=" + attr1[2]:                              "ajp_error_count",
		catalina1 + name1 + ",attr=" + attr1[3]:                              "ajp_processing_time",
		catalina1 + name1 + ",attr=" + attr1[4]:                              "ajp_request_count",
		catalina1 + name1 + ",attr=" + attr1[5]:                              "ajp_max_time",
		catalina1 + name1 + ",attr=" + attr1[6]:                              "ajp_modeler_type",
		catalina1 + name2 + ",attr=" + attr1[0]:                              "http_bytes_received",
		catalina1 + name2 + ",attr=" + attr1[1]:                              "http_bytes_sent",
		catalina1 + name2 + ",attr=" + attr1[2]:                              "http_error_count",
		catalina1 + name2 + ",attr=" + attr1[3]:                              "http_processing_time",
		catalina1 + name2 + ",attr=" + attr1[4]:                              "http_request_count",
		catalina1 + name2 + ",attr=" + attr1[5]:                              "http_max_time",
		catalina1 + name2 + ",attr=" + attr1[6]:                              "http_modeler_type",
		"java.lang:type=OperatingSystem,attr=OpenFileDescriptorCount":        "open_file_descriptor_count",
		"java.lang:type=OperatingSystem,attr=CommittedVirtualMemorySize":     "committed_memory_size",
		"java.lang:type=OperatingSystem,attr=FreePhysicalMemorySize":         "free_physical_memory_size",
		"java.lang:type=OperatingSystem,attr=ProcessCpuLoad":                 "process_cpu_load",
		"java.lang:type=OperatingSystem,attr=FreeSwapSpaceSize":              "free_swap_space_size",
		"java.lang:type=OperatingSystem,attr=TotalPhysicalMemorySize":        "total_physical_memory_size",
		"java.lang:type=OperatingSystem,attr=Name":                           "os_name",
		"java.lang:type=OperatingSystem,attr=ProcessCpuTime":                 "cpu_process_time",
		"java.lang:type=OperatingSystem,attr=MaxFileDescriptorCount":         "max_file_descriptor_count",
		"java.lang:type=OperatingSystem,attr=SystemCpuLoad":                  "system_cpu_load",
		"java.lang:type=OperatingSystem,attr=AvailableProcessors":            "available_processor_count",
		"java.lang:type=OperatingSystem,attr=Arch":                           "arch",
		"java.lang:type=OperatingSystem,attr=TotalSwapSpaceSize":             "total_swap_space_size",
		"java.lang:type=OperatingSystem,attr=Version":                        "os_version",
		"java.lang:type=Compilation,attr=TotalCompilationTime":               "total_compilation_time",
		"java.lang:type=Compilation,attr=Name":                               "compilation_name",
		"java.lang:type=Compilation,attr=CompilationTimeMonitoringSupported": "compilation_time_monitoring_supported",
		"java.lang:type=Memory,attr=HeapMemoryUsage.Committed":               "heap_memory_committed",
		"java.lang:type=Memory,attr=HeapMemoryUsage.Max":                     "heap_memory_max",
		"java.lang:type=Memory,attr=HeapMemoryUsage.Used":                    "heap_memory_used",
		"java.lang:type=Memory,attr=HeapMemoryUsage.Init":                    "heap_memory_init",
		"java.lang:type=Memory,attr=NonHeapMemoryUsage.Committed":            "non_heap_memory_committed",
		"java.lang:type=Memory,attr=NonHeapMemoryUsage.Init":                 "non_heap_memory_init",
		"java.lang:type=Memory,attr=NonHeapMemoryUsage.Max":                  "non_heap_memory_max",
		"java.lang:type=Memory,attr=NonHeapMemoryUsage.Used":                 "non_heap_memory_used",
		"java.lang:type=Memory,attr=ObjectPendingFinalizationCount":          "object_pending_finalization_count",
		"java.lang:type=Memory,attr=Verbose":                                 "memory_verbose",
		"java.lang:type=Threading,attr=ThreadCount":                          "thread_count",
		"java.lang:type=Threading,attr=TotalStartedThreadCount":              "total_started_thread_count",
		"java.lang:type=Threading,attr=DaemonThreadCount":                    "daemon_thread_count",
		"java.lang:type=Threading,attr=PeakThreadCount":                      "peak_thread_count",
		"java.lang:type=Threading,attr=CurrentThreadCpuTime":                 "current_thread_cpu_time",
		"java.lang:type=Threading,attr=CurrentThreadCpuTimeSupported":        "current_thread_cpu_time_supported",
		"java.lang:type=Threading,attr=CurrentThreadUserTime":                "current_thread_user_time",
		"java.lang:type=Threading,attr=ObjectMonitorUsageSupported":          "object_monitor_usage_supported",
		"java.lang:type=Threading,attr=SynchronizerUsageSupported":           "synchronizer_usage_supported",
		"java.lang:type=Threading,attr=ThreadAllocatedMemoryEnabled":         "thread_allocated_memory_enabled",
		"java.lang:type=Threading,attr=ThreadAllocatedMemorySupported":       "thread_allocated_memory_supported",
		"java.lang:type=Threading,attr=ThreadContentionMonitoringEnabled":    "thread_contention_monitoring_enabled",
		"java.lang:type=Threading,attr=ThreadContentionMonitoringSupported":  "thread_contention_monitoring_supported",
		"java.lang:type=Threading,attr=ThreadCpuTimeEnabled":                 "thread_cpu_time_enabled",
		"java.lang:type=Threading,attr=ThreadCpuTimeSupported":               "thread_cpu_time_supported",
		"java.lang:type=Runtime,attr=Name":                                   "runtime_name",
		"java.lang:type=Runtime,attr=Uptime":                                 "uptime",
		"java.lang:type=Runtime,attr=VmName":                                 "jvm_name",
		"java.lang:type=Runtime,attr=VmVersion":                              "jvm_version",
		"java.lang:type=Runtime,attr=StartTime":                              "jvm_start_time",
		"java.lang:type=Runtime,attr=SpecName":                               "runtime_spec_name",
		"java.lang:type=Runtime,attr=BootClassPath":                          "boo_class_path",
		"java.lang:type=Runtime,attr=BootClassPathSupported":                 "boot_class_path_supported",
		"java.lang:type=Runtime,attr=ClassPath":                              "class_path",
		"java.lang:type=Runtime,attr=LibraryPath":                            "library_path",
		"java.lang:type=Runtime,attr=ManagementSpecVersion":                  "management_spec_version",
		"java.lang:type=Runtime,attr=SpecVendor":                             "spec_vendor",
		"java.lang:type=Runtime,attr=SpecVersion":                            "spec_version",
		"java.lang:type=Runtime,attr=VmVendor":                               "vm_vendor",
		"java.lang:type=ClassLoading,attr=LoadedClassCount":                  "loaded_class_count",
		"java.lang:type=ClassLoading,attr=UnloadedClassCount":                "unloaded_class_count",
		"java.lang:type=ClassLoading,attr=TotalLoadedClassCount":             "total_loaded_class_count",
		"java.lang:type=ClassLoading,attr=Verbose":                           "classloading_verbose",
		catalina2 + port1 + ",attr=" + attr2[0]:                              port1 + "_" + "accept_count",
		catalina2 + port1 + ",attr=" + attr2[1]:                              port1 + "_" + "allow_trace",
		catalina2 + port1 + ",attr=" + attr2[2]:                              port1 + "_" + "connection_linger",
		catalina2 + port1 + ",attr=" + attr2[3]:                              port1 + "_" + "connection_timeout",
		catalina2 + port1 + ",attr=" + attr2[4]:                              port1 + "_" + "enable_lookups",
		catalina2 + port1 + ",attr=" + attr2[5]:                              port1 + "_" + "executor_name",
		catalina2 + port1 + ",attr=" + attr2[6]:                              port1 + "_" + "keep_alive_timeout",
		catalina2 + port1 + ",attr=" + attr2[7]:                              port1 + "_" + "local_port",
		catalina2 + port1 + ",attr=" + attr2[8]:                              port1 + "_" + "max_header_count",
		/*catalina2 + port1 + ",attr=" + attr2[9]:                              port1 + "_" + "max_keep_alive_requests",\\ */
		catalina2 + port1 + ",attr=" + attr2[10]:                             port1 + "_" + "max_parameter_count",
		catalina2 + port1 + ",attr=" + attr2[11]:                             port1 + "_" + "max_post_size",
		catalina2 + port1 + ",attr=" + attr2[12]:                             port1 + "_" + "max_save_post_size",
		/*catalina2 + port1 + ",attr=" + attr2[13]:                             port1 + "_" + "max_swallow_size",\\ */
		catalina2 + port1 + ",attr=" + attr2[14]:                             port1 + "_" + "max_threads",
		catalina2 + port1 + ",attr=" + attr2[15]:                             port1 + "_" + "min_spare_threads",
		catalina2 + port1 + ",attr=" + attr2[16]:                             port1 + "_" + "port",
		catalina2 + port1 + ",attr=" + attr2[17]:                             port1 + "_" + "processor_cache",
		catalina2 + port1 + ",attr=" + attr2[18]:                             port1 + "_" + "protocol",
		catalina2 + port1 + ",attr=" + attr2[19]:                             port1 + "_" + "protocol_handler_class_name",
		catalina2 + port1 + ",attr=" + attr2[20]:                             port1 + "_" + "proxy_port",
		catalina2 + port1 + ",attr=" + attr2[21]:                             port1 + "_" + "redirect_port",
		catalina2 + port1 + ",attr=" + attr2[22]:                             port1 + "_" + "scheme",
		catalina2 + port1 + ",attr=" + attr2[23]:                             port1 + "_" + "secure",
		catalina2 + port1 + ",attr=" + attr2[24]:                             port1 + "_" + "state_name",
		catalina2 + port1 + ",attr=" + attr2[25]:                             port1 + "_" + "tcp_no_delay",
		catalina2 + port1 + ",attr=" + attr2[26]:                             port1 + "_" + "thread_priority",
		catalina2 + port1 + ",attr=" + attr2[27]:                             port1 + "_" + "use_body_encoding_uri",
		catalina2 + port1 + ",attr=" + attr2[28]:                             port1 + "_" + "use_ipvhosts",
		catalina2 + port1 + ",attr=" + attr2[29]:                             port1 + "_" + "xpowered_by",
		catalina2 + port2 + ",attr=" + attr2[0]:                              port2 + "_" + "accept_count",
		catalina2 + port2 + ",attr=" + attr2[1]:                              port2 + "_" + "allow_trace",
		catalina2 + port2 + ",attr=" + attr2[2]:                              port2 + "_" + "connection_linger",
		catalina2 + port2 + ",attr=" + attr2[3]:                              port2 + "_" + "connection_timeout",
		catalina2 + port2 + ",attr=" + attr2[4]:                              port2 + "_" + "enable_lookups",
		catalina2 + port2 + ",attr=" + attr2[5]:                              port2 + "_" + "executor_name",
		catalina2 + port2 + ",attr=" + attr2[6]:                              port2 + "_" + "keep_alive_timeout",
		catalina2 + port2 + ",attr=" + attr2[7]:                              port2 + "_" + "local_port",
		catalina2 + port2 + ",attr=" + attr2[8]:                              port2 + "_" + "max_header_count",
		catalina2 + port2 + ",attr=" + attr2[9]:                              port2 + "_" + "max_keep_alive_requests",
		catalina2 + port2 + ",attr=" + attr2[10]:                             port2 + "_" + "max_parameter_count",
		catalina2 + port2 + ",attr=" + attr2[11]:                             port2 + "_" + "max_post_size",
		catalina2 + port2 + ",attr=" + attr2[12]:                             port2 + "_" + "max_save_post_size",
		catalina2 + port2 + ",attr=" + attr2[13]:                             port2 + "_" + "max_swallow_size",
		catalina2 + port2 + ",attr=" + attr2[14]:                             port2 + "_" + "max_threads",
		catalina2 + port2 + ",attr=" + attr2[15]:                             port2 + "_" + "min_spare_threads",
		catalina2 + port2 + ",attr=" + attr2[16]:                             port2 + "_" + "port",
		catalina2 + port2 + ",attr=" + attr2[17]:                             port2 + "_" + "processor_cache",
		catalina2 + port2 + ",attr=" + attr2[18]:                             port2 + "_" + "protocol",
		catalina2 + port2 + ",attr=" + attr2[19]:                             port2 + "_" + "protocol_handler_class_name",
		catalina2 + port2 + ",attr=" + attr2[20]:                             port2 + "_" + "proxy_port",
		catalina2 + port2 + ",attr=" + attr2[21]:                             port2 + "_" + "redirect_port",
		catalina2 + port2 + ",attr=" + attr2[22]:                             port2 + "_" + "scheme",
		catalina2 + port2 + ",attr=" + attr2[23]:                             port2 + "_" + "secure",
		catalina2 + port2 + ",attr=" + attr2[24]:                             port2 + "_" + "state_name",
		catalina2 + port2 + ",attr=" + attr2[25]:                             port2 + "_" + "tcp_no_delay",
		catalina2 + port2 + ",attr=" + attr2[26]:                             port2 + "_" + "thread_priority",
		catalina2 + port2 + ",attr=" + attr2[27]:                             port2 + "_" + "use_body_encoding_uri",
		catalina2 + port2 + ",attr=" + attr2[28]:                             port2 + "_" + "use_ipvhosts",
		catalina2 + port2 + ",attr=" + attr2[29]:                             port2 + "_" + "xpowered_by",
	}

}

type tomcat struct {
	id, ip, port, identification string
}

var index []string = []string{"Catalina:type=Connector,port=*", "Catalina:type=GlobalRequestProcessor,name=*", "Catalina:type=Manager,context=*,host=localhost", "java.lang:type=OperatingSystem", "java.lang:type=Compilation", "java.lang:type=Memory", "java.nio:type=BufferPool", "java.lang:type=Threading", "java.lang:type=Runtime", "java.lang:type=ClassLoading"}

//var index []string = []string{"Catalina:type=Connector,port=*" , "Catalina:type=Engine" ,"Catalina:type=Host,host=localhost" , "Catalina:type=ProtocolHandler,port=8009"}
//
