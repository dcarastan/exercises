#! /bin/bash -e

#
# Aporeto problem #1: US states file generator
#

usage() {
    cat <<EOF
$0 -- Creates a US states file

Author: Doru Carastan (doru@rocketmail.com)

Usage:  $0 [--help|-h]
        $0 --create-file=<filename> [--no-prompt] [--verbose]

EOF
}

US_STATES='Alabama
Alaska
Arizona
Arkansas
California
Colorado
Connecticut
Delaware
Florida
Georgia
Hawaii
Idaho
Illinois
Indiana
Iowa
Kansas
Kentucky
Louisiana
Maine
Maryland
Massachusetts
Michigan
Minnesota
Mississippi
Missouri
Montana
Nebraska
Nevada
New Hampshire
New Jersey
New Mexico
New York
North Carolina
North Dakota
Ohio
Oklahoma
Oregon
Pennsylvania
Rhode Island
South Carolina
South Dakota
Tennessee
Texas
Utah
Vermont
Virginia
Washington
West Virginia
Wisconsin
Wyoming'

# Set default (unset) values for command line option variables
unset opt_create_file
unset opt_no_prompt
unset opt_verbose
unset rm_flags

# Process the command line options
for option in "$@"; do
    case "${option}" in
        --create-file=*)
            opt_create_file="${option#*=}"
            shift
            ;;
        --no-prompt)
            opt_no_prompt="1"
            shift
            ;;
        --verbose)
            opt_verbose="1"
            shift
            ;;
        -h|--help)
            usage
            exit
            ;;
        --)
            # Command line option processing stops at this marker
            break
            ;;
        *)
            echo "Unknown option '${option}'" >&2
            usage >&2
            exit 2
        ;;
    esac
done

# Check manadatory options
if [ -z "${opt_create_file}" ]; then
    echo "ERROR: Specify the output file using the '--create-file' option" >&2
    usage >&2
    exit 2
fi

# If the output file exists we will have to prompt the user before removing it.
if [ -e "${opt_create_file}" ]; then
    test -n "${opt_verbose}" && echo 'File already exists'
    if [[ ${opt_no_prompt} -ne 1 ]]; then
        overwrite='n'
        while [ "${overwrite}" != 'y' ]; do
            echo -n 'File exists. Overwrite (y/n) ? '
            read overwrite
        done
    fi
    rm -- "${opt_create_file}"
    # rm returns success when the user denies its prompt to delete a R/O file
    if [ -f "${opt_create_file}" ]; then
        echo "ERROR: Failed to remove file: ${opt_create_file}" >&2
        exit 2
    else
        test -n "${opt_verbose}" && echo 'File removed'
    fi
fi

# Create the US states file
echo "${US_STATES}" > "${opt_create_file}" &&
    test -n "${opt_verbose}" && echo 'File created'

