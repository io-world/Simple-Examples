Full Screen Diff Comparison Tool

This is a tool to compare 2 images of the same dimensions and will return
whether the 2 images are identical and if not, will also return an image
which highlights the differences.

This is designed to be used in conjuction with Eggplant Functional where an
existing screenshot of a SUT is sent in along with the current screenshot to
detect if there have been any changes.

--------------------------------------------------------------------------------

Installation:

After extracting, place screen_compare/ and screen_compare.suite/ on the disk.
The helper script in screen_compare.suite needs to be updated so that
`screen_compare_bin` points to where the executable is saved.

To use the helper suite screen_compare to your own suite, go to the suite's settings
and add screen_compare.suite to the Helper Suites list.

--------------------------------------------------------------------------------

Included is a Eggplant Functional helper suite to use the full screen comparison
tool. This will take a reference screenshot as what the screen should look like,
as well as an optional mask to block out certain parts of the image which will
change (eg. dates, times, daily offers etc.); from here, the script will take a
screenshot of the current state of the SUT and call the executable.

Example usage:

// First, take a screenshot of the SUT
CaptureScreen(Name: "...")
put the result into RefScreenshotPath

// Do some tests on the SUT
...

// Then, check if the screenshot has changed
set screenCompareResult to screen_compare(
    RefScreenshotPath, /*optional*/ LiveScreenshotPath, /*optional*/ maskPath, /*optional*/ OutputImagePath
)

// parse result correctly
// result JSON is of the form
// {
//     'success': true/false,
//
//     /* in response if success == true */
//     'result': {
//         'images_match': true/false,
//         'output_filepath': string /* exists if images_match is false */
//     },
//
//     /* in response if success == false */
//     'error': string,
// }

--------------------------------------------------------------------------------


Executable usage
----------------

USAGE:
    rs-imagediff [OPTIONS] --input <input> <input>

FLAGS:
    -h, --help       Prints help information
    -V, --version    Prints version information

OPTIONS:
    -i, --input <input> <input>    the paths of the input images
        --mask <mask>              path to the mask to ignore certain parts of the image
    -o, --output <output>          the path of the output image

<input> is the file paths for the 2 input images which are to be compared. These
must be of the same dimension (width x height), filetype, channel information etc.

<output> (optional) is the filepath where the resulting image should be saved if
the 2 inputs are different.
If <output> is not specified, it will be saved to the directory from which
the shell command is called.

If the defaults are used, please make sure that the relevant folder exists on your
filesystem so that the image can be saved.

--------------------------------------------------------------------------------

This is licensed under the Apache License Version 2.0.
Please see "LICENSE.txt" for details.
Bundled third-party software may have its own license.
Please see "LICENSE-3RD-PARTY.txt for details.
