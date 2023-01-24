#!/bin/bash -xe

Image_Copy_pkg()
{
	BLD_Image_Location=${WORKSPACE}/dist
    for goos_goarch in ${pairs_without_exe}
    	do
        	fname=${BLD_Image_Location}/terraform-provider-fortiadc-internal_${goos_goarch}/${executable_name} #  terraform-provider-fortimanager_v1.6.0-SNAPSHOT-3992fc5a
      		if [ -f ${fname} ]
            then
                cp ${fname} ${PKG_dir}/${executable_rename}_${goos_goarch}
            else
                echo "${fname} not copied to ${PKG_dir}."
                exit 1
            fi
    	done
	
    for goos_goarch in ${pairs_with_exe}
    	do
        	fname=${BLD_Image_Location}/terraform-provider-fortiadc-internal_${goos_goarch}/${executable_name}.exe #  terraform-provider-fortimanager_v1.6.0-SNAPSHOT-3992fc5a
      		if [ -f ${fname} ]
            then
                cp ${fname} ${PKG_dir}/${executable_rename}_${goos_goarch}.exe
            else
                echo "${fname} not copied to ${PKG_dir}."
                exit 1
            fi
    	done
}


Md5sum()
 {
	cd ${PKG_dir}
	md5sum * | tee md5sum.txt
	sha512sum * | tee sha512sum.txt
}

File_Copy()
{
    echo "copying to nfs"
    NFS_dir=/nfs/Build_PKG/FortiAnalyzerTerraform/v${Build_Major}.00/images
    #export HISTIGNORE='*sudo -S*'
    
    if [ ! -d ${NFS_dir} ]
       then 
    	  #echo ${terraform_password} | 
          sudo -S mkdir -p ${NFS_dir}
    fi
    
    #echo ${terraform_password} | 
    sudo -S rsync -avrz ${PKG_dir} ${NFS_dir}   
	echo ${build_incr_no} ${Build_Major} ${NFS_dir}
    echo "copying to the Infosite"
    python3 /nfs/Build_PKG/jenkins_changes/change_log_4.py "FortiADCTerraform" ${Build_Major} ${Build_Minor} ${Patch} ${BN} ${Branch_Name} "artifact2infosite" ''
    echo "Files are copied to the Infosite"
}

### Main program
echo "FortiAnalyzerTerraform Build is Started"

# Fortianalyzer Internal
build_incr_no=build${BN}
PKG_dir=/home/terraform/pkg/terraform-provider-fortiadc-internal/${build_incr_no} #/home/terraform/pkg/${build_incr_no}

# if this pkg directory exists already, delete it and create it afresh
if [ -d "${PKG_dir}" ]
then
	rm -R ${PKG_dir}
fi
mkdir -p ${PKG_dir}


cd ${WORKSPACE} # /home/terraform/gitlab/terraform-fortiadc
# 1. setup workspace
export GPG_TTY=$(tty) 
#export GITHUB_TOKEN="5" #<YOUR GITHUB TOKEN>
#export GPG_FINGERPRINT=5H
export GO111MODULE=on # keep as off for internal build
export http_proxy=http://172.30.41.250:3128
export https_proxy=http://172.30.41.250:3128

#git diff
#exit 1

echo "Before Export"
echo ${PATH}
export PATH=$PATH:/usr/local/go/bin:/usr/local/bin/goreleaser
echo "After Export"
echo ${PATH}

#3. Goreleaser generate binary file
goreleaser build  --debug --snapshot --rm-dist 

commit=$(git rev-parse --short HEAD) # get HEAD shortened (7 characters)
latest_tag_commit=$(git rev-list --tags --max-count=1 | head -c7) # get latest tag. IT ISN'T NECESSARILY ASSOCIATED WITH HEAD!. this is needed because the executables are named with the latest tag
latest_tag=$(git describe --tags ${latest_tag_commit}) # (contd. from above) regardless of what is passed in "previous tag" 
executable_name=terraform-provider-fortiadc-internal_v${latest_tag}-SNAPSHOT-${commit}
executable_rename=terraform-provider-fortiadc_v${Build_Major}-build${BN}
echo ${commit}
echo ${latest_tag}
echo ${executable_name}
echo ${executable_rename}
#exit 0

Image_Copy_pkg # copy binaries to package directory
Md5sum # we still need this
File_Copy 

#echo "fortiadc-terraform Build is Completed"
# if testing DONT FORGET TO DISABLE JOB PASSED MONITOR
