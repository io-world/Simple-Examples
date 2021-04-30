import xmlrpc.client

def main():
    try:
        # Create the server proxy
        server = xmlrpc.client.ServerProxy("http://127.0.0.1:5400")

        # Start the session
        server.startsession("/Users/ranhesse/Documents/Simple_Examples/eggDrive.suite/")

        # Connect to a SUT
        server.execute ('connect (name:"TutorialSUT",port:"3389")')

        result= server.execute ('log suiteinfo()')
        print (result)

        # Launch Chrome (Where, "Chrome" is an image of its icon)
        result= server.execute ('DoubleClick "Star"')
        print (result)

        # Wait for the search box to appear
        #result= server.execute ('Click (Text: "Search Google or type a URL", WaitFor: 10)')
        #print (result)

        # Go to the Eggplant website
        #result= server.execute ('TypeText("https://www.eggplantsoftware.com" & returnKey)')
        #print (result)

        # Wait for the banner
        #result= server.execute('WaitFor 10.0', '(Text: "Let\'s rid the world of bad software")')
        #print (result)

        # Quit Chrome
        #result= server.execute ('TypeText(altKey, f4)')
        #print (result)
    except Exception as exc:
        print("An exception occurred: {}".format(exc))

    finally:
        try:
            # End the session
            server.endsession ("")
        except Exception as exc:
            print("An exception occurred: {}".format(exc))

if __name__ == '__main__':
    main()