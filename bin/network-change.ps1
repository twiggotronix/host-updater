# Get the active network adapters (those that are "Up")
$activeAdapters = Get-NetAdapter | Where-Object { $_.Status -eq "Up" }

# Check if there is an active Ethernet (wired) connection
$ethernetAdapter = $activeAdapters | Where-Object { $_.MediaType -eq "802.3" -and $_.InterfaceDescription -notmatch "Virtual|Hyper-V|VMware|VirtualBox" }

# Check if there is an active Wi-Fi (wireless) connection
$wifiAdapter = $activeAdapters | Where-Object { $_.MediaType -in @("802.11", "Native 802.11") }

if ($ethernetAdapter) {
    # If an Ethernet connection is active, run the specified command for wired connection
    Start-Process "C:\projects\host-updater\bin\host-updater.exe" -ArgumentList "update"
    Write-Output "Running host-updater for wired connection."
}
elseif ($wifiAdapter) {
    # If a Wi-Fi connection is active, run the specified command for wireless connection
    Start-Process "C:\projects\host-updater\bin\host-updater.exe" -ArgumentList "update -w"
    Write-Output "Running host-updater for Wi-Fi connection."
}
else {
    Write-Output "No active network connection detected."
}
# Start-Sleep -Seconds 10