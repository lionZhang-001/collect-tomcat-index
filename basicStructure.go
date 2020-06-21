/*
@Time : 2020/6/19 20:57
@Author : zhangyongyue
@File : basicStructure
@Software: GoLand
*/
package main

type tomcat struct {
	id, ip, port, identification string
}

var index []string = []string{"Catalina:type=Connector,port=*" ,"Catalina:type=GlobalRequestProcessor,name=*", "Catalina:type=Manager,context=*,host=localhost", "java.lang:type=OperatingSystem", "java.lang:type=Compilation", "java.lang:type=Memory", "java.nio:type=BufferPool", "java.lang:type=Threading", "java.lang:type=Runtime", "java.lang:type=ClassLoading"}
//var index []string = []string{"Catalina:type=Connector,port=*" , "Catalina:type=Engine" ,"Catalina:type=Host,host=localhost" , "Catalina:type=ProtocolHandler,port=8009"}

var index2Mysql = map[string]string{"Catalina:type=GlobalRequestProcessor,name=\"ajp-bio-8009\",attr=bytesReceived": "ajp_bytes_received",
	"Catalina:type=GlobalRequestProcessor,name=\"ajp-bio-8009\",attr=bytesSent":       "ajp_bytes_sent",
	"Catalina:type=GlobalRequestProcessor,name=\"ajp-bio-8009\",attr=errorCount":      "ajp_error_count",
	"Catalina:type=GlobalRequestProcessor,name=\"ajp-bio-8009\",attr=processingTime":  "ajp_processing_time",
	"Catalina:type=GlobalRequestProcessor,name=\"ajp-bio-8009\",attr=requestCount":    "ajp_request_count",
	"Catalina:type=GlobalRequestProcessor,name=\"http-bio-8888\",attr=bytesReceived":  "http_bytes_received",
	"Catalina:type=GlobalRequestProcessor,name=\"http-bio-8888\",attr=bytesSent":      "http_bytes_sent",
	"Catalina:type=GlobalRequestProcessor,name=\"http-bio-8888\",attr=requestCount":   "http_request_count",
	"Catalina:type=GlobalRequestProcessor,name=\"http-bio-8888\",attr=processingTime": "http_processing_time",
	"Catalina:type=GlobalRequestProcessor,name=\"http-bio-8888\",attr=errorCount":     "http_error_count",
	"java.lang:type=OperatingSystem,attr=OpenFileDescriptorCount":                     "open_file_descriptor_count",
	"java.lang:type=OperatingSystem,attr=CommittedVirtualMemorySize":                  "committed_memory_size",
	"java.lang:type=OperatingSystem,attr=FreePhysicalMemorySize":                      "free_physical_memory_size",
	"java.lang:type=OperatingSystem,attr=ProcessCpuLoad":                              "process_cpu_load",
	"java.lang:type=OperatingSystem,attr=FreeSwapSpaceSize":                           "free_swap_space_size",
	"java.lang:type=OperatingSystem,attr=TotalPhysicalMemorySize":                     "total_physical_memory_size",
	"java.lang:type=OperatingSystem,attr=Name":                                        "os_name",
	"java.lang:type=OperatingSystem,attr=ProcessCpuTime":                              "cpu_process_time",
	"java.lang:type=OperatingSystem,attr=MaxFileDescriptorCount":                      "max_file_descriptor_count",
	"java.lang:type=OperatingSystem,attr=SystemCpuLoad":                               "system_cpu_load",
	"java.lang:type=OperatingSystem,attr=AvailableProcessors":                         "available_processor_count",
	"java.lang:type=Compilation,attr=TotalCompilationTime":                            "total_compilation_time",
	"java.lang:type=Compilation,attr=Name":                                            "compilation_name",
	"java.lang:type=Memory,attr=HeapMemoryUsage.Committed":                            "heap_memory_committed",
	"java.lang:type=Memory,attr=HeapMemoryUsage.Max":                                  "heap_memory_max",
	"java.lang:type=Memory,attr=HeapMemoryUsage.Used":                                 "heap_memory_used",
	"java.lang:type=Memory,attr=HeapMemoryUsage.Init":                                 "heap_memory_init",
	"java.lang:type=Memory,attr=NonHeapMemoryUsage.Committed":                         "non_heap_memory_committed",
	"java.lang:type=Memory,attr=NonHeapMemoryUsage.Init":                              "non_heap_memory_init",
	"java.lang:type=Memory,attr=NonHeapMemoryUsage.Max":                               "non_heap_memory_max",
	"java.lang:type=Memory,attr=NonHeapMemoryUsage.Used":                              "non_heap_memory_used",
	"java.lang:type=Threading,attr=ThreadCount":                                       "thread_count",
	"java.lang:type=Threading,attr=TotalStartedThreadCount":                           "total_started_thread_count",
	"java.lang:type=Threading,attr=DaemonThreadCount":                                 "daemon_thread_count",
	"java.lang:type=Threading,attr=PeakThreadCount":                                   "peak_thread_count",
	"java.lang:type=Runtime,attr=Uptime":                                              "uptime",
	"java.lang:type=Runtime,attr=VmName":                                              "jvm_name",
	"java.lang:type=Runtime,attr=VmVersion":                                           "jvm_version",
	"java.lang:type=Runtime,attr=StartTime":                                           "jvm_start_time",
	"java.lang:type=ClassLoading,attr=LoadedClassCount":                               "loaded_class_count",
	"java.lang:type=ClassLoading,attr=UnloadedClassCount":                             "unloaded_class_count",
	"java.lang:type=ClassLoading,attr=TotalLoadedClassCount":                          "total_loaded_class_count",
}
