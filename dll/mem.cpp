#include <iostream>
#include <Windows.h>

#include "mem.h"

/**
 * Set the page guard flag on a block of memory to catch read/writes
 * @param address - address ptr
 * @param length - mem block length
 */
auto PageGuardMemory(void* address, const size_t length) -> void
{
    DWORD oldProtect;
    MEMORY_BASIC_INFORMATION mbi;

    VirtualQuery(static_cast<const void*>(address), &mbi, sizeof(MEMORY_BASIC_INFORMATION));
    VirtualProtect(address, length, mbi.Protect | PAGE_GUARD, &oldProtect);
}

/**
 * Attempt to remove the page guard flag from a memory block
 * @param address - address ptr
 * @param length - mem block length
 */
auto UnPageGuardMemory(void* address, const size_t length) -> void
{
    DWORD oldProtect;
    MEMORY_BASIC_INFORMATION mbi;

    VirtualQuery(static_cast<const void*>(address), &mbi, sizeof(MEMORY_BASIC_INFORMATION));
    VirtualProtect(address, length, mbi.Protect & ~PAGE_GUARD, &oldProtect);
}

/**
 * Check if a given address is readable and writable
 * @param address - address ptr
 */
auto CheckAddrReadable(void* address) -> bool
{
    MEMORY_BASIC_INFORMATION mbi;

    VirtualQuery(static_cast<const void*>(address), &mbi, sizeof(MEMORY_BASIC_INFORMATION));
    if (mbi.Protect & ~PAGE_GUARD) {
        return false;
    } else if (mbi.Protect & ~PAGE_NOACCESS) {
        return false;
    }
    return false;
}

/**
 * Callback function to catch access violations or page guards so we don't crash the game
 * @param ep
 * @param address
 * @return
 */
auto CALLBACK VectoredExceptionHandler(_EXCEPTION_POINTERS* ep, void* address) -> LONG
{
    if (ep->ExceptionRecord->ExceptionCode == EXCEPTION_GUARD_PAGE)
    {
        if (ep->ExceptionRecord->ExceptionInformation[1] == reinterpret_cast<ULONG_PTR>(address))
        {
            std::cout << "Memory page guarded at address " << std::hex << ep->ExceptionRecord->ExceptionAddress <<
                      ", address accessed: " << ep->ExceptionRecord->ExceptionInformation[1] << std::dec << std::endl;
        }
        ep->ContextRecord->EFlags |= 0x100ui32;
        return EXCEPTION_CONTINUE_EXECUTION;
    }
    else if(ep->ExceptionRecord->ExceptionCode == EXCEPTION_ACCESS_VIOLATION)
    {
        ep->ContextRecord->EFlags |= 0x01ui32;
        return EXCEPTION_CONTINUE_EXECUTION;
    }
    return EXCEPTION_CONTINUE_SEARCH;
}