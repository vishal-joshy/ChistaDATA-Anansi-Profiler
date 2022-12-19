package data_handlers

import "github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"

// GetDataHandler using configuration it selects the best possible IDataHandler
func GetDataHandler(config *stucts.CliConfig) IDataHandler {
	if len(config.FilePaths) > 0 {
		return InitLocalFilesDataHandler(config.FilePaths, config.DatabaseName, config.DatabaseVersion)
	}
	return nil
}
