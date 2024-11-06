function Test-LDAPPorts {
    [CmdletBinding()]
    param(
        [string] $ServerName,
        [int] $Port
    )
    if ($ServerName -and $Port -ne 0) {
        try {
            $LDAP = "LDAP://" + $ServerName + ':' + $Port
            $Connection = [ADSI]($LDAP)
            $Connection.Close()
            return $true
        } catch {
            if ($_.Exception.ToString() -match "The server is not operational") {
                Write-Warning "Can't open $ServerName`:$Port."
            } elseif ($_.Exception.ToString() -match "The user name or password is incorrect") {
                Write-Warning "Current user ($Env:USERNAME) doesn't seem to have access to to LDAP on port $Server`:$Port"
            } else {
                Write-Warning -Message $_
            }
        }
        return $False
    }
}
Function Test-LDAP {
    [CmdletBinding()]
    param (
        [alias('Server', 'IpAddress')][Parameter(Mandatory = $True)][string[]]$ComputerName,
        [int] $GCPortLDAP = 3268,
        [int] $GCPortLDAPSSL = 3269,
        [int] $PortLDAP = 10389,
        [int] $PortLDAPS = 636
    )
    # Checks for ServerName - Makes sure to convert IPAddress to DNS
    foreach ($Computer in $ComputerName) {
        [Array] $ADServerFQDN = (Resolve-DnsName -Name $Computer -ErrorAction SilentlyContinue)
        if ($ADServerFQDN) {
            if ($ADServerFQDN.NameHost) {
                $ServerName = $ADServerFQDN[0].NameHost
            } else {
                [Array] $ADServerFQDN = (Resolve-DnsName -Name $Computer -ErrorAction SilentlyContinue)
                $FilterName = $ADServerFQDN | Where-Object { $_.QueryType -eq 'A' }
                $ServerName = $FilterName[0].Name
            }
        } else {
            $ServerName = ''
        }

        $GlobalCatalogSSL = Test-LDAPPorts -ServerName $ServerName -Port $GCPortLDAPSSL
        $GlobalCatalogNonSSL = Test-LDAPPorts -ServerName $ServerName -Port $GCPortLDAP
        $ConnectionLDAPS = Test-LDAPPorts -ServerName $ServerName -Port $PortLDAPS
        $ConnectionLDAP = Test-LDAPPorts -ServerName $ServerName -Port $PortLDAP

        $PortsThatWork = @(
            if ($GlobalCatalogNonSSL) { $GCPortLDAP }
            if ($GlobalCatalogSSL) { $GCPortLDAPSSL }
            if ($ConnectionLDAP) { $PortLDAP }
            if ($ConnectionLDAPS) { $PortLDAPS }
        ) | Sort-Object
        [pscustomobject]@{
            Computer           = $Computer
            ComputerFQDN       = $ServerName
            GlobalCatalogLDAP  = $GlobalCatalogNonSSL
            GlobalCatalogLDAPS = $GlobalCatalogSSL
            LDAP               = $ConnectionLDAP
            LDAPS              = $ConnectionLDAPS
            AvailablePorts     = $PortsThatWork -join ','
        }
    }
}

Function Test-LDAPConnection {
    [CmdletBinding()]
               
    # Parameters used in this function
    Param
    (
        [Parameter(Position=0, Mandatory = $True, HelpMessage="Provide domain controllers names, example DC01", ValueFromPipeline = $true)] 
        $DCs,
  
        [Parameter(Position=1, Mandatory = $False, HelpMessage="Provide port number for LDAP", ValueFromPipeline = $true)] 
        $Port = "636"
    ) 
  
    $ErrorActionPreference = "Stop"
    $Results = @()
    Try{ 
        Import-Module ActiveDirectory -ErrorAction Stop
    }
    Catch{
        $_.Exception.Message
        Break
    } 
         
    ForEach($DC in $DCs){
        $DC =$DC.trim()
        Write-Verbose "Processing $DC"
        Try{
            $DCName = (Get-ADDomainController -Identity $DC).hostname
        }
        Catch{
            $_.Exception.Message
            Continue
        }
  
        If($DCName -ne $Null){  
            Try{
                $Connection = [adsi]"LDAP://$($DCName):$Port"
            }
            Catch{
                $ExcMessage = $_.Exception.Message
                throw "Error: Failed to make LDAP connection. Exception: $ExcMessage"
            }
  
            If ($Connection.Path) {
                $Object = New-Object PSObject -Property ([ordered]@{ 
                       
                    DC                = $DC
                    Port              = $Port
                    Path              = $Connection.Path
                })
  
                $Results += $Object
            }         
        }
    }
  
    If($Results){
        Return $Results
    }
 }